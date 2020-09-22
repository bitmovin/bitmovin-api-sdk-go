package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioAacCustomdataAPI communicates with '/encoding/configurations/audio/aac/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioAacCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioAacCustomdataAPI constructor for EncodingConfigurationsAudioAacCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioAacCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioAacCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioAacCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioAacCustomdataAPIWithClient constructor for EncodingConfigurationsAudioAacCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioAacCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioAacCustomdataAPI {
	a := &EncodingConfigurationsAudioAacCustomdataAPI{apiClient: apiClient}
	return a
}

// Get AAC Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioAacCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/aac/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
