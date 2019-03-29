package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsHttpsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsHttpsCustomdataApi
}

func NewEncodingInputsHttpsApi(configs ...func(*common.ApiClient)) (*EncodingInputsHttpsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsHttpsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsHttpsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsHttpsApi) Create(httpsInput model.HttpsInput) (*model.HttpsInput, error) {
    payload := model.HttpsInput(httpsInput)
    
    err := api.apiClient.Post("/encoding/inputs/https", &payload)
    return &payload, err
}
func (api *EncodingInputsHttpsApi) Delete(inputId string) (*model.HttpsInput, error) {
    var resp *model.HttpsInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Delete("/encoding/inputs/https/{input_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsHttpsApi) List(queryParams ...func(*query.HttpsInputListQueryParams)) (*pagination.HttpsInputsListPagination, error) {
    queryParameters := &query.HttpsInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.HttpsInputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/inputs/https", &resp, reqParams)
    return resp, err
}
func (api *EncodingInputsHttpsApi) Get(inputId string) (*model.HttpsInput, error) {
    var resp *model.HttpsInput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/https/{input_id}", &resp, reqParams)
    return resp, err
}
