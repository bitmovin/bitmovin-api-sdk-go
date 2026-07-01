package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoH265v2API communicates with '/encoding/configurations/video/h265v2' endpoints
type EncodingConfigurationsVideoH265v2API struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoH265v2API constructor for EncodingConfigurationsVideoH265v2API that takes options as argument
func NewEncodingConfigurationsVideoH265v2API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH265v2API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoH265v2APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH265v2APIWithClient constructor for EncodingConfigurationsVideoH265v2API that takes an APIClient as argument
func NewEncodingConfigurationsVideoH265v2APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH265v2API {
	a := &EncodingConfigurationsVideoH265v2API{apiClient: apiClient}
	return a
}

// Create H265 V2 Codec Configuration
func (api *EncodingConfigurationsVideoH265v2API) Create(h265V2VideoConfiguration model.H265V2VideoConfiguration) (*model.H265V2VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.H265V2VideoConfiguration
	err := api.apiClient.Post("/encoding/configurations/video/h265v2", &h265V2VideoConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete H265 V2 Codec Configuration
func (api *EncodingConfigurationsVideoH265v2API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/video/h265v2/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get H265 V2 Codec Configuration Details
func (api *EncodingConfigurationsVideoH265v2API) Get(configurationId string) (*model.H265V2VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.H265V2VideoConfiguration
	err := api.apiClient.Get("/encoding/configurations/video/h265v2/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List H265 V2 Codec Configurations
func (api *EncodingConfigurationsVideoH265v2API) List(queryParams ...func(*EncodingConfigurationsVideoH265v2APIListQueryParams)) (*pagination.H265V2VideoConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsVideoH265v2APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.H265V2VideoConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/video/h265v2", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsVideoH265v2APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoH265v2APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoH265v2APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
