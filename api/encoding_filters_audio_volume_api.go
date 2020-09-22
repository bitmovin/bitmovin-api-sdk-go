package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersAudioVolumeAPI communicates with '/encoding/filters/audio-volume' endpoints
type EncodingFiltersAudioVolumeAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/audio-volume/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersAudioVolumeCustomdataAPI
}

// NewEncodingFiltersAudioVolumeAPI constructor for EncodingFiltersAudioVolumeAPI that takes options as argument
func NewEncodingFiltersAudioVolumeAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAudioVolumeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersAudioVolumeAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAudioVolumeAPIWithClient constructor for EncodingFiltersAudioVolumeAPI that takes an APIClient as argument
func NewEncodingFiltersAudioVolumeAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAudioVolumeAPI {
	a := &EncodingFiltersAudioVolumeAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersAudioVolumeCustomdataAPIWithClient(apiClient)

	return a
}

// Create Audio Volume Filter
func (api *EncodingFiltersAudioVolumeAPI) Create(audioVolumeFilter model.AudioVolumeFilter) (*model.AudioVolumeFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AudioVolumeFilter
	err := api.apiClient.Post("/encoding/filters/audio-volume", &audioVolumeFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Audio Volume Filter
func (api *EncodingFiltersAudioVolumeAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/audio-volume/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Audio Volume Filter Details
func (api *EncodingFiltersAudioVolumeAPI) Get(filterId string) (*model.AudioVolumeFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.AudioVolumeFilter
	err := api.apiClient.Get("/encoding/filters/audio-volume/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Audio Volume Filters
func (api *EncodingFiltersAudioVolumeAPI) List(queryParams ...func(*EncodingFiltersAudioVolumeAPIListQueryParams)) (*pagination.AudioVolumeFiltersListPagination, error) {
	queryParameters := &EncodingFiltersAudioVolumeAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AudioVolumeFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/audio-volume", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersAudioVolumeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersAudioVolumeAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersAudioVolumeAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
