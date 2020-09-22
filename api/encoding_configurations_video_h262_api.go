package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoH262API communicates with '/encoding/configurations/video/h262' endpoints
type EncodingConfigurationsVideoH262API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/video/h262/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsVideoH262CustomdataAPI
}

// NewEncodingConfigurationsVideoH262API constructor for EncodingConfigurationsVideoH262API that takes options as argument
func NewEncodingConfigurationsVideoH262API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH262API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoH262APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH262APIWithClient constructor for EncodingConfigurationsVideoH262API that takes an APIClient as argument
func NewEncodingConfigurationsVideoH262APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH262API {
	a := &EncodingConfigurationsVideoH262API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsVideoH262CustomdataAPIWithClient(apiClient)

	return a
}

// Create H262 Codec Configuration
func (api *EncodingConfigurationsVideoH262API) Create(h262VideoConfiguration model.H262VideoConfiguration) (*model.H262VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.H262VideoConfiguration
	err := api.apiClient.Post("/encoding/configurations/video/h262", &h262VideoConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete H262 Codec Configuration
func (api *EncodingConfigurationsVideoH262API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/video/h262/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get H262 Codec Configuration Details
func (api *EncodingConfigurationsVideoH262API) Get(configurationId string) (*model.H262VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.H262VideoConfiguration
	err := api.apiClient.Get("/encoding/configurations/video/h262/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List H262 Codec Configurations
func (api *EncodingConfigurationsVideoH262API) List(queryParams ...func(*EncodingConfigurationsVideoH262APIListQueryParams)) (*pagination.H262VideoConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsVideoH262APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.H262VideoConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/video/h262", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsVideoH262APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoH262APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoH262APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
