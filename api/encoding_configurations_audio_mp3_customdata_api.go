package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioMp3CustomdataAPI communicates with '/encoding/configurations/audio/mp3/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioMp3CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioMp3CustomdataAPI constructor for EncodingConfigurationsAudioMp3CustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioMp3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioMp3CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioMp3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioMp3CustomdataAPIWithClient constructor for EncodingConfigurationsAudioMp3CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioMp3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioMp3CustomdataAPI {
	a := &EncodingConfigurationsAudioMp3CustomdataAPI{apiClient: apiClient}
	return a
}

// Get MP3 Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioMp3CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/mp3/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
