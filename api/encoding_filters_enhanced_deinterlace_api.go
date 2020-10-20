package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersEnhancedDeinterlaceAPI communicates with '/encoding/filters/enhanced-deinterlace' endpoints
type EncodingFiltersEnhancedDeinterlaceAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/enhanced-deinterlace/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersEnhancedDeinterlaceCustomdataAPI
}

// NewEncodingFiltersEnhancedDeinterlaceAPI constructor for EncodingFiltersEnhancedDeinterlaceAPI that takes options as argument
func NewEncodingFiltersEnhancedDeinterlaceAPI(options ...apiclient.APIClientOption) (*EncodingFiltersEnhancedDeinterlaceAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersEnhancedDeinterlaceAPIWithClient(apiClient), nil
}

// NewEncodingFiltersEnhancedDeinterlaceAPIWithClient constructor for EncodingFiltersEnhancedDeinterlaceAPI that takes an APIClient as argument
func NewEncodingFiltersEnhancedDeinterlaceAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersEnhancedDeinterlaceAPI {
	a := &EncodingFiltersEnhancedDeinterlaceAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersEnhancedDeinterlaceCustomdataAPIWithClient(apiClient)

	return a
}

// Create Enhanced Deinterlace Filter
func (api *EncodingFiltersEnhancedDeinterlaceAPI) Create(enhancedDeinterlaceFilter model.EnhancedDeinterlaceFilter) (*model.EnhancedDeinterlaceFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.EnhancedDeinterlaceFilter
	err := api.apiClient.Post("/encoding/filters/enhanced-deinterlace", &enhancedDeinterlaceFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Enhanced Deinterlace Filter
func (api *EncodingFiltersEnhancedDeinterlaceAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/enhanced-deinterlace/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Enhanced Deinterlace Filter Details
func (api *EncodingFiltersEnhancedDeinterlaceAPI) Get(filterId string) (*model.EnhancedDeinterlaceFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.EnhancedDeinterlaceFilter
	err := api.apiClient.Get("/encoding/filters/enhanced-deinterlace/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Enhanced Deinterlace Filters
func (api *EncodingFiltersEnhancedDeinterlaceAPI) List(queryParams ...func(*EncodingFiltersEnhancedDeinterlaceAPIListQueryParams)) (*pagination.EnhancedDeinterlaceFiltersListPagination, error) {
	queryParameters := &EncodingFiltersEnhancedDeinterlaceAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EnhancedDeinterlaceFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/enhanced-deinterlace", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersEnhancedDeinterlaceAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersEnhancedDeinterlaceAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersEnhancedDeinterlaceAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
