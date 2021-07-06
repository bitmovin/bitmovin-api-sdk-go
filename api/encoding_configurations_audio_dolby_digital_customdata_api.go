package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioDolbyDigitalCustomdataAPI communicates with '/encoding/configurations/audio/dolby-digital/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioDolbyDigitalCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioDolbyDigitalCustomdataAPI constructor for EncodingConfigurationsAudioDolbyDigitalCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioDolbyDigitalCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDolbyDigitalCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDolbyDigitalCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDolbyDigitalCustomdataAPIWithClient constructor for EncodingConfigurationsAudioDolbyDigitalCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDolbyDigitalCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDolbyDigitalCustomdataAPI {
	a := &EncodingConfigurationsAudioDolbyDigitalCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Dolby Digital Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioDolbyDigitalCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-digital/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
