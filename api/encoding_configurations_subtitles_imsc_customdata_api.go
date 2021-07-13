package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsSubtitlesImscCustomdataAPI communicates with '/encoding/configurations/subtitles/imsc/{configuration_id}/customData' endpoints
type EncodingConfigurationsSubtitlesImscCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsSubtitlesImscCustomdataAPI constructor for EncodingConfigurationsSubtitlesImscCustomdataAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesImscCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesImscCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsSubtitlesImscCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesImscCustomdataAPIWithClient constructor for EncodingConfigurationsSubtitlesImscCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesImscCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesImscCustomdataAPI {
	a := &EncodingConfigurationsSubtitlesImscCustomdataAPI{apiClient: apiClient}
	return a
}

// Get IMSC subtitle configuration custom data
func (api *EncodingConfigurationsSubtitlesImscCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/subtitles/imsc/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
