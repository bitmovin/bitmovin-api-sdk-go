package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsAudioPcmAPI communicates with '/encoding/configurations/audio/pcm' endpoints
type EncodingConfigurationsAudioPcmAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/audio/pcm/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsAudioPcmCustomdataAPI
}

// NewEncodingConfigurationsAudioPcmAPI constructor for EncodingConfigurationsAudioPcmAPI that takes options as argument
func NewEncodingConfigurationsAudioPcmAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioPcmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioPcmAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioPcmAPIWithClient constructor for EncodingConfigurationsAudioPcmAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioPcmAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioPcmAPI {
	a := &EncodingConfigurationsAudioPcmAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsAudioPcmCustomdataAPIWithClient(apiClient)

	return a
}

// Create PCM Codec Configuration
func (api *EncodingConfigurationsAudioPcmAPI) Create(pcmAudioConfiguration model.PcmAudioConfiguration) (*model.PcmAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.PcmAudioConfiguration
	err := api.apiClient.Post("/encoding/configurations/audio/pcm", &pcmAudioConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete PCM Codec Configuration
func (api *EncodingConfigurationsAudioPcmAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/audio/pcm/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get PCM Codec Configuration Details
func (api *EncodingConfigurationsAudioPcmAPI) Get(configurationId string) (*model.PcmAudioConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.PcmAudioConfiguration
	err := api.apiClient.Get("/encoding/configurations/audio/pcm/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List PCM Configurations
func (api *EncodingConfigurationsAudioPcmAPI) List(queryParams ...func(*EncodingConfigurationsAudioPcmAPIListQueryParams)) (*pagination.PcmAudioConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsAudioPcmAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.PcmAudioConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/audio/pcm", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsAudioPcmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsAudioPcmAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsAudioPcmAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
