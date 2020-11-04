package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioAc3CustomdataAPI communicates with '/encoding/configurations/audio/ac3/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioAc3CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioAc3CustomdataAPI constructor for EncodingConfigurationsAudioAc3CustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioAc3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioAc3CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioAc3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioAc3CustomdataAPIWithClient constructor for EncodingConfigurationsAudioAc3CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioAc3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioAc3CustomdataAPI {
	a := &EncodingConfigurationsAudioAc3CustomdataAPI{apiClient: apiClient}
	return a
}

// Get AC3 Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioAc3CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/ac3/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
