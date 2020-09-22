package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioHeAacV2CustomdataAPI communicates with '/encoding/configurations/audio/he-aac-v2/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioHeAacV2CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioHeAacV2CustomdataAPI constructor for EncodingConfigurationsAudioHeAacV2CustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioHeAacV2CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioHeAacV2CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioHeAacV2CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioHeAacV2CustomdataAPIWithClient constructor for EncodingConfigurationsAudioHeAacV2CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioHeAacV2CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioHeAacV2CustomdataAPI {
	a := &EncodingConfigurationsAudioHeAacV2CustomdataAPI{apiClient: apiClient}
	return a
}

// Get HE-AAC v2 Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioHeAacV2CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v2/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
