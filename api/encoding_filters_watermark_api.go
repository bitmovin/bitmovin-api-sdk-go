package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersWatermarkAPI communicates with '/encoding/filters/watermark' endpoints
type EncodingFiltersWatermarkAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/watermark/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersWatermarkCustomdataAPI
}

// NewEncodingFiltersWatermarkAPI constructor for EncodingFiltersWatermarkAPI that takes options as argument
func NewEncodingFiltersWatermarkAPI(options ...apiclient.APIClientOption) (*EncodingFiltersWatermarkAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersWatermarkAPIWithClient(apiClient), nil
}

// NewEncodingFiltersWatermarkAPIWithClient constructor for EncodingFiltersWatermarkAPI that takes an APIClient as argument
func NewEncodingFiltersWatermarkAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersWatermarkAPI {
	a := &EncodingFiltersWatermarkAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersWatermarkCustomdataAPIWithClient(apiClient)

	return a
}

// Create Watermark Filter
func (api *EncodingFiltersWatermarkAPI) Create(watermarkFilter model.WatermarkFilter) (*model.WatermarkFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.WatermarkFilter
	err := api.apiClient.Post("/encoding/filters/watermark", &watermarkFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Watermark Filter
func (api *EncodingFiltersWatermarkAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/watermark/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Watermark Filter Details
func (api *EncodingFiltersWatermarkAPI) Get(filterId string) (*model.WatermarkFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.WatermarkFilter
	err := api.apiClient.Get("/encoding/filters/watermark/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Watermark Filters
func (api *EncodingFiltersWatermarkAPI) List(queryParams ...func(*EncodingFiltersWatermarkAPIListQueryParams)) (*pagination.WatermarkFiltersListPagination, error) {
	queryParameters := &EncodingFiltersWatermarkAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WatermarkFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/watermark", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersWatermarkAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersWatermarkAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersWatermarkAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
