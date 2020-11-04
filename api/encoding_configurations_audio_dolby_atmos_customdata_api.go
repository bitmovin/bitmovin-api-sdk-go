package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioDolbyAtmosCustomdataAPI communicates with '/encoding/configurations/audio/dolby-atmos/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioDolbyAtmosCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioDolbyAtmosCustomdataAPI constructor for EncodingConfigurationsAudioDolbyAtmosCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioDolbyAtmosCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDolbyAtmosCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDolbyAtmosCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDolbyAtmosCustomdataAPIWithClient constructor for EncodingConfigurationsAudioDolbyAtmosCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDolbyAtmosCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDolbyAtmosCustomdataAPI {
	a := &EncodingConfigurationsAudioDolbyAtmosCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Dolby Atmos Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioDolbyAtmosCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-atmos/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
