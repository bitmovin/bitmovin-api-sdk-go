package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioDtsCustomdataAPI communicates with '/encoding/configurations/audio/dts/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioDtsCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioDtsCustomdataAPI constructor for EncodingConfigurationsAudioDtsCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioDtsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDtsCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDtsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDtsCustomdataAPIWithClient constructor for EncodingConfigurationsAudioDtsCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDtsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDtsCustomdataAPI {
	a := &EncodingConfigurationsAudioDtsCustomdataAPI{apiClient: apiClient}
	return a
}

// Get DTS Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioDtsCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/dts/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
