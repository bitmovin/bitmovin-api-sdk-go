package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioPassthroughAPI communicates with '/encoding/configurations/audio/passthrough' endpoints
type EncodingConfigurationsAudioPassthroughAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioPassthroughAPI constructor for EncodingConfigurationsAudioPassthroughAPI that takes options as argument
func NewEncodingConfigurationsAudioPassthroughAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioPassthroughAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioPassthroughAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioPassthroughAPIWithClient constructor for EncodingConfigurationsAudioPassthroughAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioPassthroughAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioPassthroughAPI {
	a := &EncodingConfigurationsAudioPassthroughAPI{apiClient: apiClient}
	return a
}

// Create Audio Passthrough Configuration
func (api *EncodingConfigurationsAudioPassthroughAPI) Create(passthroughAudioConfiguration model.PassthroughAudioConfiguration) (*model.PassthroughAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.PassthroughAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/passthrough", &passthroughAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Audio Passthrough Codec Configuration
func (api *EncodingConfigurationsAudioPassthroughAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/passthrough/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Audio Passthrough Configuration Details
func (api *EncodingConfigurationsAudioPassthroughAPI) Get(configurationId string) (*model.PassthroughAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.PassthroughAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/passthrough/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Audio Passthrough Configurations
func (api *EncodingConfigurationsAudioPassthroughAPI) List(queryParams ...func(*EncodingConfigurationsAudioPassthroughAPIListQueryParams)) (*pagination.PassthroughAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioPassthroughAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PassthroughAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/passthrough", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioPassthroughAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioPassthroughAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioPassthroughAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
