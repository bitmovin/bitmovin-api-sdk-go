package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersEnhancedWatermarkAPI communicates with '/encoding/filters/enhanced-watermark' endpoints
type EncodingFiltersEnhancedWatermarkAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/enhanced-watermark/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersEnhancedWatermarkCustomdataAPI
}

// NewEncodingFiltersEnhancedWatermarkAPI constructor for EncodingFiltersEnhancedWatermarkAPI that takes options as argument
func NewEncodingFiltersEnhancedWatermarkAPI(options ...apiclient.APIClientOption) (*EncodingFiltersEnhancedWatermarkAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersEnhancedWatermarkAPIWithClient(apiClient), nil
}

// NewEncodingFiltersEnhancedWatermarkAPIWithClient constructor for EncodingFiltersEnhancedWatermarkAPI that takes an APIClient as argument
func NewEncodingFiltersEnhancedWatermarkAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersEnhancedWatermarkAPI {
	a := &EncodingFiltersEnhancedWatermarkAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersEnhancedWatermarkCustomdataAPIWithClient(apiClient)

	return a
}

// Create Enhanced Watermark Filter
func (api *EncodingFiltersEnhancedWatermarkAPI) Create(enhancedWatermarkFilter model.EnhancedWatermarkFilter) (*model.EnhancedWatermarkFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.EnhancedWatermarkFilter
	err := api.apiClient.Post("/encoding/filters/enhanced-watermark", &enhancedWatermarkFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Enhanced Watermark Filter
func (api *EncodingFiltersEnhancedWatermarkAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/enhanced-watermark/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Enhanced Watermark Filter Details
func (api *EncodingFiltersEnhancedWatermarkAPI) Get(filterId string) (*model.EnhancedWatermarkFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.EnhancedWatermarkFilter
	err := api.apiClient.Get("/encoding/filters/enhanced-watermark/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Enhanced Watermark Filters
func (api *EncodingFiltersEnhancedWatermarkAPI) List(queryParams ...func(*EncodingFiltersEnhancedWatermarkAPIListQueryParams)) (*pagination.EnhancedWatermarkFiltersListPagination, error) {
	queryParameters := &EncodingFiltersEnhancedWatermarkAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EnhancedWatermarkFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/enhanced-watermark", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersEnhancedWatermarkAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersEnhancedWatermarkAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersEnhancedWatermarkAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
