package apiclient

import (
	. "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	. "github.com/bitmovin/bitmovin-api-sdk-go/model"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// APIClient tha will execute all HTTP requests to the Bitmovin API
type APIClient struct {
	BaseURL     string
	APIKey      string
	TenantOrgId string
	Logger      Logger

	restClient *RestClient
}

// APIClientOption is a shorthand for the initializing function
type APIClientOption func(*APIClient)

// WithBaseURL to configure the baseURL of the APIClient
func WithBaseURL(baseURL string) APIClientOption {
	return func(apiClient *APIClient) {
		apiClient.BaseURL = baseURL
	}
}

// WithAPIKey to configure the Logger of the APIClient
func WithAPIKey(apiKey string) APIClientOption {
	return func(apiClient *APIClient) {
		apiClient.APIKey = apiKey
	}
}

// WithTenantOrgId to configure the Logger of the APIClient
func WithTenantOrgId(tenantOrgId string) APIClientOption {
	return func(apiClient *APIClient) {
		apiClient.TenantOrgId = tenantOrgId
	}
}

// WithLogger to configure the Logger of the APIClient
func WithLogger(logger Logger) APIClientOption {
	return func(apiClient *APIClient) {
		apiClient.Logger = logger
	}
}

type PathParams map[string]interface{}

const QueryParamTagName = "query"
const DefaultAPIBaseURL = "https://api.bitmovin.com/v1"
const ContentTypeJson = "application/json"
const APIClientVersion = "1.222.0"
const APIClientName = "bitmovin-api-sdk-go"
const NoAPIKeyErrorMsg = "there was no api key provided"

type QueryParams interface {
	Params() map[string]string
}

type RequestParams struct {
	PathParams  PathParams
	QueryParams QueryParams
}

func NewAPIClient(options ...APIClientOption) (*APIClient, error) {
	apiClient := APIClient{
		BaseURL: DefaultAPIBaseURL,
		Logger:  &StdOutLogger{},
	}

	for _, option := range options {
		option(&apiClient)
	}

	if apiClient.APIKey == "" {
		return nil, fmt.Errorf(NoAPIKeyErrorMsg)
	}

	apiClient.restClient = NewRestClient(&Configuration{})
	return &apiClient, nil
}

func GetParamsMap(u interface{}) map[string]string {
	val := reflect.ValueOf(u)
	if val.Kind() != reflect.Ptr {
		panic("Parameter of GetParamsMap has to be a pointer")
	}

	val = reflect.ValueOf(u).Elem()
	t := val.Type()

	queryParamMap := make(map[string]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		typeField := t.Field(i)
		tag := typeField.Tag.Get(QueryParamTagName)

		valueField := val.Field(i)

		switch valueField.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := valueField.Addr().Elem().Int()
			if value > 0 {
				queryParamMap[tag] = strconv.FormatInt(value, 10)
			}
		case reflect.String:
			value := valueField.Addr().Elem().String()
			if len(value) > 0 {
				queryParamMap[tag] = value
			}
		case reflect.Struct:
			emptyTime := time.Time{}.Unix()
			value := valueField.Interface()
			date, ok := value.(Date)
			if ok && time.Time(date).Unix() != emptyTime {
				queryParamMap[tag] = date.String()
			}
			datetime, ok := value.(DateTime)
			if ok && time.Time(datetime).Unix() != emptyTime {
				queryParamMap[tag] = datetime.StringUTC()
			}
		}
	}
	return queryParamMap
}

func (apiClient *APIClient) prepareURL(relURL string, reqParams ...func(reqParams *RequestParams)) (string, error) {
	u, err := url.Parse(apiClient.BaseURL)
	if err != nil {
		panic(err)
	}

	requestParams := &RequestParams{
		PathParams: make(map[string]interface{}),
	}
	for _, reqParam := range reqParams {
		reqParam(requestParams)
	}

	regex := regexp.MustCompile(`{(.*?)}`)
	found := regex.FindAllStringSubmatch(relURL, -1)

	for k, v := range requestParams.PathParams {
		stringVal := ""
		if timeParam, isTime := v.(time.Time); isTime {
			stringVal = timeParam.Format(time.RFC3339)
		} else if stringParam, isString := v.(string); isString {
			stringVal = stringParam
		} else {
			return "", fmt.Errorf("path param can not be of type %s", reflect.TypeOf(v).String())
		}

		for _, f := range found {
			if k == f[1] {
				relURL = strings.Replace(relURL, "{"+f[1]+"}", stringVal, -1)
			}
		}
	}

	found = regex.FindAllStringSubmatch(relURL, -1)
	if len(found) > 0 {
		return "", fmt.Errorf("there are still some placeholders left in the URL. "+
			"Please make sure to pass the correct values to this function to replace all placeholders. url=%s", relURL)
	}

	if requestParams.QueryParams != nil {
		queryParamMap := requestParams.QueryParams.Params()
		var queryParams []string
		queryParamString := ""
		for k, v := range queryParamMap {
			queryParam := fmt.Sprintf("%s=%s", k, v)
			queryParams = append(queryParams, queryParam)
		}
		queryParamString = strings.Join(queryParams, "&")
		if len(queryParamString) > 0 {
			u.RawQuery = queryParamString
		}
	}

	u.Path = path.Join(u.Path, relURL)
	return u.String(), nil
}

func (apiClient *APIClient) createHeaders() http.Header {
	headers := http.Header{}
	headers.Add("X-Api-Key", apiClient.APIKey)
	if apiClient.TenantOrgId != "" {
		headers.Add(
			"X-Tenant-Org-Id", apiClient.TenantOrgId,
		)
	}
	headers.Add("X-Api-Client", APIClientName)
	headers.Add("X-Api-Client-Version", APIClientVersion)
	headers.Add("Content-Type", ContentTypeJson)

	return headers
}

func (apiClient *APIClient) Get(relURL string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodGet, relURL, nil, responseModel, requestParams...)
}

func (apiClient *APIClient) Post(relURL string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodPost, relURL, requestModel, responseModel, requestParams...)
}

func (apiClient *APIClient) Patch(relURL string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodPatch, relURL, requestModel, responseModel, requestParams...)
}

func (apiClient *APIClient) Put(relURL string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodPut, relURL, requestModel, responseModel, requestParams...)
}

func (apiClient *APIClient) Delete(relURL string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodDelete, relURL, nil, responseModel, requestParams...)
}

func (apiClient *APIClient) logRequest(reqURL string, method string, headers http.Header, rawBody []byte) {
	apiClient.Logger.LogRequest(RequestLogInfo{
		URL:     reqURL,
		Method:  method,
		Headers: headers,
		RawBody: rawBody,
	})
}

func (apiClient *APIClient) logResponse(reqURL string, method string, headers http.Header, statusCode int, rawBody []byte) {
	apiClient.Logger.LogResponse(ResponseLogInfo{
		URL:        reqURL,
		Method:     method,
		Headers:    headers,
		RawBody:    rawBody,
		StatusCode: statusCode,
	})
}

func buildBitmovinError(envelope GenericResponseEnvelope, method string, reqURL string, requestString string, responseString string, httpStatusCode *int) BitmovinError {
	var bitErr BitmovinError
	json.Unmarshal(envelope.Data, &bitErr)
	bitErr.RawResponse = responseString
	bitErr.HttpStatusCode = httpStatusCode
	bitErr.ShortMessage = bitErr.Message
	bitErr.Message = buildErrorMessage(bitErr, method, reqURL, requestString, responseString, httpStatusCode)
	bitErr.RequestId = envelope.RequestId
	return bitErr
}

func buildShortErrorMessage(method string, reqURL string, requestString string, code *int) string {
	var sb strings.Builder
	sb.WriteString(buildRequest(method, reqURL, requestString))
	sb.WriteString("\n")
	sb.WriteString(buildResponse(code, ""))
	return sb.String()
}

func buildSimpleErrorMessage(method string, reqURL string, requestString string, errorMessage string) string {
	var sb strings.Builder
	sb.WriteString("Request failed: ")
	sb.WriteString(errorMessage)
	sb.WriteString("\n")
	sb.WriteString(buildRequest(method, reqURL, requestString))
	sb.WriteString("\n")
	return sb.String()
}

func buildErrorMessage(bitmovinError BitmovinError, method string, reqURL string, requestString string, responseString string, code *int) string {
	var sb strings.Builder
	if bitmovinError.Message != "" {
		sb.WriteString(bitmovinError.Message)
		sb.WriteString("\n")
	}

	if bitmovinError.DeveloperMessage != "" {
		sb.WriteString("developerMessage: " + bitmovinError.DeveloperMessage)
		sb.WriteString("\n")
	}

	if bitmovinError.ErrorCode != nil {
		sb.WriteString(fmt.Sprintf("errorCode: %d", Int32Value(bitmovinError.ErrorCode)))
		sb.WriteString("\n")
	}

	if len(bitmovinError.Details) > 0 {
		sb.WriteString(buildDetails(bitmovinError.Details))
		sb.WriteString("\n")
	}

	if len(bitmovinError.Links) > 0 {
		sb.WriteString(buildLinks(bitmovinError.Links))
		sb.WriteString("\n")
	}

	sb.WriteString(buildRequest(method, reqURL, requestString))
	sb.WriteString("\n")
	sb.WriteString(buildResponse(code, responseString))

	return sb.String()
}

func buildDetails(details []Message) string {
	var sb strings.Builder
	var detail Message
	sb.WriteString("details:")
	for _, detail = range details {
		if detail.Id != nil {
			sb.WriteString("\n")
			sb.WriteString("  - id: " + *detail.Id)
		}

		if detail.Date != nil {
			sb.WriteString("\n")
			sb.WriteString("    date: ")
			sb.WriteString(detail.Date.StringUTC())
		}

		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("    type: %v", detail.Type))
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("    text: %s", StringValue(detail.Text)))

		if detail.Field != nil {
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("    field: %s", StringValue(detail.Field)))
		}
	}

	return sb.String()
}

func buildLinks(links []Link) string {
	var sb strings.Builder
	var link Link
	sb.WriteString("links:")
	for _, link = range links {
		sb.WriteString("\n")
		sb.WriteString("  " + *link.Title + ": " + *link.Href)
	}
	return sb.String()
}

func buildRequest(method string, reqURL string, requestString string) string {
	var sb strings.Builder

	sb.WriteString("request:")
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("  method: %s", method))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("  url: %s", reqURL))
	if requestString != "" {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("  body: %s", requestString))
	}

	return sb.String()
}

func buildResponse(httpStatusCode *int, responseString string) string {
	var sb strings.Builder
	sb.WriteString("response:")
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("  httpStatusCode: %d", IntValue(httpStatusCode)))
	if responseString != "" {
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("  body: %s", responseString))
	}

	return sb.String()
}

func (apiClient *APIClient) request(reqMethod string, relURL string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {

	reqURL, err := apiClient.prepareURL(relURL, requestParams...)
	if err != nil {
		return err
	}

	headers := apiClient.createHeaders()
	if err != nil {
		return err
	}

	var resp *http.Response
	var requestString string
	switch reqMethod {
	case http.MethodPost:
		serialized, _ := json.Marshal(requestModel)
		apiClient.logRequest(reqURL, reqMethod, headers, serialized)
		resp, err = apiClient.restClient.Post(reqURL, serialized, headers)
		requestString = string(serialized)
	case http.MethodPatch:
		serialized, _ := json.Marshal(requestModel)
		apiClient.logRequest(reqURL, reqMethod, headers, serialized)
		resp, err = apiClient.restClient.Patch(reqURL, serialized, headers)
		requestString = string(serialized)
	case http.MethodPut:
		serialized, _ := json.Marshal(requestModel)
		apiClient.logRequest(reqURL, reqMethod, headers, serialized)
		resp, err = apiClient.restClient.Put(reqURL, serialized, headers)
		requestString = string(serialized)
	case http.MethodGet:
		apiClient.logRequest(reqURL, reqMethod, headers, nil)
		resp, err = apiClient.restClient.Get(reqURL, headers)
	case http.MethodDelete:
		apiClient.logRequest(reqURL, reqMethod, headers, nil)
		resp, err = apiClient.restClient.Delete(reqURL, headers)
	}

	if err != nil {
		return buildRequestFailedErrorMessage("", resp, reqMethod, reqURL, requestString, err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return buildRequestFailedErrorMessage(string(respBody), resp, reqMethod, reqURL, requestString, err)
	}

	apiClient.logResponse(reqURL, reqMethod, resp.Header, resp.StatusCode, respBody)

	isSuccessResponse := resp.StatusCode >= 200 && resp.StatusCode <= 299

	if len(respBody) != 0 {
		var envelope GenericResponseEnvelope
		err = json.Unmarshal(respBody, &envelope)

		if isSuccessResponse {
			convertToModel(envelope, &responseModel)
		} else {
			bitmovinError := BitmovinError{}
			if err == nil {
				bitmovinError = buildBitmovinError(envelope, reqMethod, reqURL, requestString, string(respBody), &resp.StatusCode)
			}
			return bitmovinError
		}
	}

	if !isSuccessResponse {
		return buildRequestFailedErrorMessage(string(respBody), resp, reqMethod, reqURL, requestString, err)
	}

	return nil
}

func buildRequestFailedErrorMessage(respBody string, resp *http.Response, reqMethod string, reqURL string, requestString string, err error) error {
	bitmovinError := BitmovinError{}
	if err == nil {
		bitmovinError.RawResponse = respBody
		bitmovinError.HttpStatusCode = &resp.StatusCode
		bitmovinError.Message = buildShortErrorMessage(reqMethod, reqURL, requestString, &resp.StatusCode)
	} else {
		bitmovinError := BitmovinError{}
		bitmovinError.Message = buildSimpleErrorMessage(reqMethod, reqURL, requestString, err.Error())
		return bitmovinError
	}
	return bitmovinError
}

func convertToModel(envelope GenericResponseEnvelope, out interface{}) {
	var successData ResultWrapper
	json.Unmarshal(envelope.Data, &successData)
	json.Unmarshal(successData.Result, &out)
}
