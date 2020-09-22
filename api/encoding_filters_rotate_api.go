package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersRotateAPI communicates with '/encoding/filters/rotate' endpoints
type EncodingFiltersRotateAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/rotate/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersRotateCustomdataAPI
}

// NewEncodingFiltersRotateAPI constructor for EncodingFiltersRotateAPI that takes options as argument
func NewEncodingFiltersRotateAPI(options ...apiclient.APIClientOption) (*EncodingFiltersRotateAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersRotateAPIWithClient(apiClient), nil
}

// NewEncodingFiltersRotateAPIWithClient constructor for EncodingFiltersRotateAPI that takes an APIClient as argument
func NewEncodingFiltersRotateAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersRotateAPI {
	a := &EncodingFiltersRotateAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersRotateCustomdataAPIWithClient(apiClient)

	return a
}

// Create Rotate Filter
func (api *EncodingFiltersRotateAPI) Create(rotateFilter model.RotateFilter) (*model.RotateFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.RotateFilter
	err := api.apiClient.Post("/encoding/filters/rotate", &rotateFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Rotate Filter
func (api *EncodingFiltersRotateAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/rotate/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Rotate Filter Details
func (api *EncodingFiltersRotateAPI) Get(filterId string) (*model.RotateFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.RotateFilter
	err := api.apiClient.Get("/encoding/filters/rotate/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Rotate Filters
func (api *EncodingFiltersRotateAPI) List(queryParams ...func(*EncodingFiltersRotateAPIListQueryParams)) (*pagination.RotateFiltersListPagination, error) {
	queryParameters := &EncodingFiltersRotateAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.RotateFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/rotate", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersRotateAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersRotateAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersRotateAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
