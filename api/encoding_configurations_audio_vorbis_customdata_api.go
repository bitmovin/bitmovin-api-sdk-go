package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioVorbisCustomdataAPI communicates with '/encoding/configurations/audio/vorbis/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioVorbisCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioVorbisCustomdataAPI constructor for EncodingConfigurationsAudioVorbisCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioVorbisCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioVorbisCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioVorbisCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioVorbisCustomdataAPIWithClient constructor for EncodingConfigurationsAudioVorbisCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioVorbisCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioVorbisCustomdataAPI {
	a := &EncodingConfigurationsAudioVorbisCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Vorbis Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioVorbisCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/vorbis/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
