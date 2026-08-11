package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersDolbyLoudnessAPI communicates with '/encoding/filters/dolby-loudness' endpoints
type EncodingFiltersDolbyLoudnessAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/dolby-loudness/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersDolbyLoudnessCustomdataAPI
}

// NewEncodingFiltersDolbyLoudnessAPI constructor for EncodingFiltersDolbyLoudnessAPI that takes options as argument
func NewEncodingFiltersDolbyLoudnessAPI(options ...apiclient.APIClientOption) (*EncodingFiltersDolbyLoudnessAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersDolbyLoudnessAPIWithClient(apiClient), nil
}

// NewEncodingFiltersDolbyLoudnessAPIWithClient constructor for EncodingFiltersDolbyLoudnessAPI that takes an APIClient as argument
func NewEncodingFiltersDolbyLoudnessAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersDolbyLoudnessAPI {
	a := &EncodingFiltersDolbyLoudnessAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersDolbyLoudnessCustomdataAPIWithClient(apiClient)

	return a
}

// Create Dolby Loudness Filter
func (api *EncodingFiltersDolbyLoudnessAPI) Create(dolbyLoudnessFilter model.DolbyLoudnessFilter) (*model.DolbyLoudnessFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.DolbyLoudnessFilter
	err := api.apiClient.Post("/encoding/filters/dolby-loudness", &dolbyLoudnessFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Dolby Loudness Filter
func (api *EncodingFiltersDolbyLoudnessAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/dolby-loudness/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Dolby Loudness Filter details
func (api *EncodingFiltersDolbyLoudnessAPI) Get(filterId string) (*model.DolbyLoudnessFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.DolbyLoudnessFilter
	err := api.apiClient.Get("/encoding/filters/dolby-loudness/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Dolby Loudness Filters
func (api *EncodingFiltersDolbyLoudnessAPI) List(queryParams ...func(*EncodingFiltersDolbyLoudnessAPIListQueryParams)) (*pagination.DolbyLoudnessFiltersListPagination, error) {
	queryParameters := &EncodingFiltersDolbyLoudnessAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DolbyLoudnessFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/dolby-loudness", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersDolbyLoudnessAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersDolbyLoudnessAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersDolbyLoudnessAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
