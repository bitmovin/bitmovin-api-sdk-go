package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoAv1API communicates with '/encoding/configurations/video/av1' endpoints
type EncodingConfigurationsVideoAv1API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/video/av1/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsVideoAv1CustomdataAPI
}

// NewEncodingConfigurationsVideoAv1API constructor for EncodingConfigurationsVideoAv1API that takes options as argument
func NewEncodingConfigurationsVideoAv1API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoAv1API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoAv1APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoAv1APIWithClient constructor for EncodingConfigurationsVideoAv1API that takes an APIClient as argument
func NewEncodingConfigurationsVideoAv1APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoAv1API {
	a := &EncodingConfigurationsVideoAv1API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsVideoAv1CustomdataAPIWithClient(apiClient)

	return a
}

// Create AV1 Codec Configuration
func (api *EncodingConfigurationsVideoAv1API) Create(av1VideoConfiguration model.Av1VideoConfiguration) (*model.Av1VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Av1VideoConfiguration
	err := api.apiClient.Post("/encoding/configurations/video/av1", &av1VideoConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AV1 Codec Configuration
func (api *EncodingConfigurationsVideoAv1API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/video/av1/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AV1 Codec Configuration Details
func (api *EncodingConfigurationsVideoAv1API) Get(configurationId string) (*model.Av1VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.Av1VideoConfiguration
	err := api.apiClient.Get("/encoding/configurations/video/av1/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AV1 Codec Configurations
func (api *EncodingConfigurationsVideoAv1API) List(queryParams ...func(*EncodingConfigurationsVideoAv1APIListQueryParams)) (*pagination.Av1VideoConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsVideoAv1APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Av1VideoConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/video/av1", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsVideoAv1APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoAv1APIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoAv1APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
