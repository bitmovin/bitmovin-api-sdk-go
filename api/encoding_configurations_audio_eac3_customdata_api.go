package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioEac3CustomdataAPI communicates with '/encoding/configurations/audio/eac3/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioEac3CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioEac3CustomdataAPI constructor for EncodingConfigurationsAudioEac3CustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioEac3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioEac3CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioEac3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioEac3CustomdataAPIWithClient constructor for EncodingConfigurationsAudioEac3CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioEac3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioEac3CustomdataAPI {
	a := &EncodingConfigurationsAudioEac3CustomdataAPI{apiClient: apiClient}
	return a
}

// Get E-AC3 Codec Configuration Custom Data.  Deprecation notice: use Dolby Digital Plus instead. For more information check out our tutorial here: https://bitmovin.com/docs/encoding/tutorials/how-to-create-dolby-digital-plus-encodings
func (api *EncodingConfigurationsAudioEac3CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/eac3/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
