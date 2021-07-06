package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI communicates with '/encoding/configurations/audio/dolby-digital-plus/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI constructor for EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPIWithClient constructor for EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI {
	a := &EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Dolby Digital Plus Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioDolbyDigitalPlusCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/dolby-digital-plus/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
