package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioAacAPI communicates with '/encoding/configurations/audio/aac' endpoints
type EncodingConfigurationsAudioAacAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/aac/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioAacCustomdataAPI
}

// NewEncodingConfigurationsAudioAacAPI constructor for EncodingConfigurationsAudioAacAPI that takes options as argument
func NewEncodingConfigurationsAudioAacAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioAacAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioAacAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioAacAPIWithClient constructor for EncodingConfigurationsAudioAacAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioAacAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioAacAPI {
	a := &EncodingConfigurationsAudioAacAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioAacCustomdataAPIWithClient(apiClient)

	return a
}

// Create AAC Codec Configuration
func (api *EncodingConfigurationsAudioAacAPI) Create(aacAudioConfiguration model.AacAudioConfiguration) (*model.AacAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AacAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/aac", &aacAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AAC Codec Configuration
func (api *EncodingConfigurationsAudioAacAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/aac/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AAC Codec Configuration Details
func (api *EncodingConfigurationsAudioAacAPI) Get(configurationId string) (*model.AacAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.AacAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/aac/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AAC Configurations
func (api *EncodingConfigurationsAudioAacAPI) List(queryParams ...func(*EncodingConfigurationsAudioAacAPIListQueryParams)) (*pagination.AacAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioAacAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AacAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/aac", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioAacAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioAacAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioAacAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
