package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoH264API communicates with '/encoding/configurations/video/h264' endpoints
type EncodingConfigurationsVideoH264API struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/video/h264/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsVideoH264CustomdataAPI
}

// NewEncodingConfigurationsVideoH264API constructor for EncodingConfigurationsVideoH264API that takes options as argument
func NewEncodingConfigurationsVideoH264API(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH264API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoH264APIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH264APIWithClient constructor for EncodingConfigurationsVideoH264API that takes an APIClient as argument
func NewEncodingConfigurationsVideoH264APIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH264API {
	a := &EncodingConfigurationsVideoH264API{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsVideoH264CustomdataAPIWithClient(apiClient)

	return a
}

// Create H264/AVC Codec Configuration
func (api *EncodingConfigurationsVideoH264API) Create(h264VideoConfiguration model.H264VideoConfiguration) (*model.H264VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.H264VideoConfiguration
	err := api.apiClient.Post("/encoding/configurations/video/h264", &h264VideoConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete H264/AVC Codec Configuration
func (api *EncodingConfigurationsVideoH264API) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/video/h264/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get H264/AVC Codec Configuration Details
func (api *EncodingConfigurationsVideoH264API) Get(configurationId string) (*model.H264VideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.H264VideoConfiguration
	err := api.apiClient.Get("/encoding/configurations/video/h264/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List H264/AVC Codec Configurations
func (api *EncodingConfigurationsVideoH264API) List(queryParams ...func(*EncodingConfigurationsVideoH264APIListQueryParams)) (*pagination.H264VideoConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsVideoH264APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.H264VideoConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/video/h264", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsVideoH264APIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoH264APIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoH264APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
