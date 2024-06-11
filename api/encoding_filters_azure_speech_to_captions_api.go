package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersAzureSpeechToCaptionsAPI communicates with '/encoding/filters/azure-speech-to-captions' endpoints
type EncodingFiltersAzureSpeechToCaptionsAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/azure-speech-to-captions/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersAzureSpeechToCaptionsCustomdataAPI
}

// NewEncodingFiltersAzureSpeechToCaptionsAPI constructor for EncodingFiltersAzureSpeechToCaptionsAPI that takes options as argument
func NewEncodingFiltersAzureSpeechToCaptionsAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAzureSpeechToCaptionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersAzureSpeechToCaptionsAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAzureSpeechToCaptionsAPIWithClient constructor for EncodingFiltersAzureSpeechToCaptionsAPI that takes an APIClient as argument
func NewEncodingFiltersAzureSpeechToCaptionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAzureSpeechToCaptionsAPI {
	a := &EncodingFiltersAzureSpeechToCaptionsAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersAzureSpeechToCaptionsCustomdataAPIWithClient(apiClient)

	return a
}

// Create Azure Speech to captions Filter
// This filter uses the Azure Speech Services - Speech to captions feature to transcribe an audio stream in real-time.  - Required for this filter is an Azure account with an Azure Speech service resource. - Your own Azure speech service subscription key is required.  - This filter only works for Live Encodings. - This filter transforms an audio input stream into a captions stream. - The filter must be applied to a stream that has an audio input. - The stream&#39;s codec configuration must be a subtitle or caption format (WebVtt, TTML, ..).
func (api *EncodingFiltersAzureSpeechToCaptionsAPI) Create(azureSpeechToCaptionsFilter model.AzureSpeechToCaptionsFilter) (*model.AzureSpeechToCaptionsFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AzureSpeechToCaptionsFilter
	err := api.apiClient.Post("/encoding/filters/azure-speech-to-captions", &azureSpeechToCaptionsFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Azure Speech to captions Filter
func (api *EncodingFiltersAzureSpeechToCaptionsAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/azure-speech-to-captions/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Azure Speech to captions Filter details
func (api *EncodingFiltersAzureSpeechToCaptionsAPI) Get(filterId string) (*model.AzureSpeechToCaptionsFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.AzureSpeechToCaptionsFilter
	err := api.apiClient.Get("/encoding/filters/azure-speech-to-captions/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Azure Speech to captions Filters
func (api *EncodingFiltersAzureSpeechToCaptionsAPI) List(queryParams ...func(*EncodingFiltersAzureSpeechToCaptionsAPIListQueryParams)) (*pagination.AzureSpeechToCaptionsFiltersListPagination, error) {
	queryParameters := &EncodingFiltersAzureSpeechToCaptionsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AzureSpeechToCaptionsFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/azure-speech-to-captions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersAzureSpeechToCaptionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersAzureSpeechToCaptionsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersAzureSpeechToCaptionsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
