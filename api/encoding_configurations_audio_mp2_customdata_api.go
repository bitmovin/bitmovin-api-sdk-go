package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioMp2CustomdataAPI communicates with '/encoding/configurations/audio/mp2/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioMp2CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioMp2CustomdataAPI constructor for EncodingConfigurationsAudioMp2CustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioMp2CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioMp2CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioMp2CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioMp2CustomdataAPIWithClient constructor for EncodingConfigurationsAudioMp2CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioMp2CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioMp2CustomdataAPI {
	a := &EncodingConfigurationsAudioMp2CustomdataAPI{apiClient: apiClient}
	return a
}

// Get MP2 Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioMp2CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/mp2/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
