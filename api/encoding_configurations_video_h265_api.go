package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoH265API communicates with '/encoding/configurations/video/h265' endpoints
type EncodingConfigurationsVideoH265API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/video/h265/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsVideoH265CustomdataAPI
}

// NewEncodingConfigurationsVideoH265API constructor for EncodingConfigurationsVideoH265API that takes options as argument
func NewEncodingConfigurationsVideoH265API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH265API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoH265APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH265APIWithClient constructor for EncodingConfigurationsVideoH265API that takes an APIClient as argument
func NewEncodingConfigurationsVideoH265APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH265API {
	a := &EncodingConfigurationsVideoH265API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsVideoH265CustomdataAPIWithClient(apiClient)

	return a
}

// Create H265/HEVC Codec Configuration
func (api *EncodingConfigurationsVideoH265API) Create(h265VideoConfiguration model.H265VideoConfiguration) (*model.H265VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.H265VideoConfiguration
	err := api.apiClient.Post("/encoding/configurations/video/h265", &h265VideoConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete H265/HEVC Codec Configuration
func (api *EncodingConfigurationsVideoH265API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/video/h265/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get H265/HEVC Codec Configuration Details
func (api *EncodingConfigurationsVideoH265API) Get(configurationId string) (*model.H265VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.H265VideoConfiguration
	err := api.apiClient.Get("/encoding/configurations/video/h265/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List H265/HEVC Codec Configurations
func (api *EncodingConfigurationsVideoH265API) List(queryParams ...func(*EncodingConfigurationsVideoH265APIListQueryParams)) (*pagination.H265VideoConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsVideoH265APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.H265VideoConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/video/h265", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsVideoH265APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoH265APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoH265APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
