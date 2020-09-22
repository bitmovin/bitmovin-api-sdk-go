package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersInterlaceAPI communicates with '/encoding/filters/interlace' endpoints
type EncodingFiltersInterlaceAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/interlace/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersInterlaceCustomdataAPI
}

// NewEncodingFiltersInterlaceAPI constructor for EncodingFiltersInterlaceAPI that takes options as argument
func NewEncodingFiltersInterlaceAPI(options ...apiclient.APIClientOption) (*EncodingFiltersInterlaceAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersInterlaceAPIWithClient(apiClient), nil
}

// NewEncodingFiltersInterlaceAPIWithClient constructor for EncodingFiltersInterlaceAPI that takes an APIClient as argument
func NewEncodingFiltersInterlaceAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersInterlaceAPI {
	a := &EncodingFiltersInterlaceAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersInterlaceCustomdataAPIWithClient(apiClient)

	return a
}

// Create Interlace Filter
func (api *EncodingFiltersInterlaceAPI) Create(interlaceFilter model.InterlaceFilter) (*model.InterlaceFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.InterlaceFilter
	err := api.apiClient.Post("/encoding/filters/interlace", &interlaceFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Interlace Filter
func (api *EncodingFiltersInterlaceAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/interlace/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Interlace Filter Details
func (api *EncodingFiltersInterlaceAPI) Get(filterId string) (*model.InterlaceFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.InterlaceFilter
	err := api.apiClient.Get("/encoding/filters/interlace/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Interlace Filters
func (api *EncodingFiltersInterlaceAPI) List(queryParams ...func(*EncodingFiltersInterlaceAPIListQueryParams)) (*pagination.InterlaceFiltersListPagination, error) {
	queryParameters := &EncodingFiltersInterlaceAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.InterlaceFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/interlace", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersInterlaceAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersInterlaceAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersInterlaceAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
