package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersEbuR128SinglePassAPI communicates with '/encoding/filters/ebu-r128-single-pass' endpoints
type EncodingFiltersEbuR128SinglePassAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/filters/ebu-r128-single-pass/{filter_id}/customData' endpoints
	Customdata *EncodingFiltersEbuR128SinglePassCustomdataAPI
}

// NewEncodingFiltersEbuR128SinglePassAPI constructor for EncodingFiltersEbuR128SinglePassAPI that takes options as argument
func NewEncodingFiltersEbuR128SinglePassAPI(options ...apiclient.APIClientOption) (*EncodingFiltersEbuR128SinglePassAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersEbuR128SinglePassAPIWithClient(apiClient), nil
}

// NewEncodingFiltersEbuR128SinglePassAPIWithClient constructor for EncodingFiltersEbuR128SinglePassAPI that takes an APIClient as argument
func NewEncodingFiltersEbuR128SinglePassAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersEbuR128SinglePassAPI {
	a := &EncodingFiltersEbuR128SinglePassAPI{apiClient: apiClient}
	a.Customdata = NewEncodingFiltersEbuR128SinglePassCustomdataAPIWithClient(apiClient)

	return a
}

// Create EBU R128 Single Pass Filter
func (api *EncodingFiltersEbuR128SinglePassAPI) Create(ebuR128SinglePassFilter model.EbuR128SinglePassFilter) (*model.EbuR128SinglePassFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.EbuR128SinglePassFilter
	err := api.apiClient.Post("/encoding/filters/ebu-r128-single-pass", &ebuR128SinglePassFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete EBU R128 Single Pass Filter
func (api *EncodingFiltersEbuR128SinglePassAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/filters/ebu-r128-single-pass/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get EBU R128 Single Pass Filter Details
func (api *EncodingFiltersEbuR128SinglePassAPI) Get(filterId string) (*model.EbuR128SinglePassFilter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.EbuR128SinglePassFilter
	err := api.apiClient.Get("/encoding/filters/ebu-r128-single-pass/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List EBU R128 Single Pass Filters
func (api *EncodingFiltersEbuR128SinglePassAPI) List(queryParams ...func(*EncodingFiltersEbuR128SinglePassAPIListQueryParams)) (*pagination.EbuR128SinglePassFiltersListPagination, error) {
	queryParameters := &EncodingFiltersEbuR128SinglePassAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EbuR128SinglePassFiltersListPagination
	err := api.apiClient.Get("/encoding/filters/ebu-r128-single-pass", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersEbuR128SinglePassAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersEbuR128SinglePassAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersEbuR128SinglePassAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
