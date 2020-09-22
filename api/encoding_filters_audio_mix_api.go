package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersAudioMixAPI communicates with '/encoding/filters/audio-mix' endpoints
type EncodingFiltersAudioMixAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/audio-mix/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersAudioMixCustomdataAPI
}

// NewEncodingFiltersAudioMixAPI constructor for EncodingFiltersAudioMixAPI that takes options as argument
func NewEncodingFiltersAudioMixAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAudioMixAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersAudioMixAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAudioMixAPIWithClient constructor for EncodingFiltersAudioMixAPI that takes an APIClient as argument
func NewEncodingFiltersAudioMixAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAudioMixAPI {
	a := &EncodingFiltersAudioMixAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersAudioMixCustomdataAPIWithClient(apiClient)

	return a
}

// Create Audio Mix Filter
func (api *EncodingFiltersAudioMixAPI) Create(audioMixFilter model.AudioMixFilter) (*model.AudioMixFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AudioMixFilter
	err := api.apiClient.Post("/encoding/filters/audio-mix", &audioMixFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Audio Mix Filter
func (api *EncodingFiltersAudioMixAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/audio-mix/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Audio Mix Filter Details
func (api *EncodingFiltersAudioMixAPI) Get(filterId string) (*model.AudioMixFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.AudioMixFilter
	err := api.apiClient.Get("/encoding/filters/audio-mix/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Audio Mix Filters
func (api *EncodingFiltersAudioMixAPI) List(queryParams ...func(*EncodingFiltersAudioMixAPIListQueryParams)) (*pagination.AudioMixFiltersListPagination, error) {
	queryParameters := &EncodingFiltersAudioMixAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AudioMixFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/audio-mix", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersAudioMixAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersAudioMixAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersAudioMixAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
