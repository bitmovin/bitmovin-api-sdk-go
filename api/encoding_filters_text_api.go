package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersTextAPI communicates with '/encoding/filters/text' endpoints
type EncodingFiltersTextAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/text/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersTextCustomdataAPI
}

// NewEncodingFiltersTextAPI constructor for EncodingFiltersTextAPI that takes options as argument
func NewEncodingFiltersTextAPI(options ...apiclient.APIClientOption) (*EncodingFiltersTextAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersTextAPIWithClient(apiClient), nil
}

// NewEncodingFiltersTextAPIWithClient constructor for EncodingFiltersTextAPI that takes an APIClient as argument
func NewEncodingFiltersTextAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersTextAPI {
	a := &EncodingFiltersTextAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersTextCustomdataAPIWithClient(apiClient)

	return a
}

// Create Text Filter
func (api *EncodingFiltersTextAPI) Create(textFilter model.TextFilter) (*model.TextFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.TextFilter
	err := api.apiClient.Post("/encoding/filters/text", &textFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Text Filter
func (api *EncodingFiltersTextAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/text/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Text Filter Details
func (api *EncodingFiltersTextAPI) Get(filterId string) (*model.TextFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.TextFilter
	err := api.apiClient.Get("/encoding/filters/text/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Text Filters
func (api *EncodingFiltersTextAPI) List(queryParams ...func(*EncodingFiltersTextAPIListQueryParams)) (*pagination.TextFiltersListPagination, error) {
	queryParameters := &EncodingFiltersTextAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.TextFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/text", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersTextAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersTextAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersTextAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
