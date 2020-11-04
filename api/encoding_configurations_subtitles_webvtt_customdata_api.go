package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsSubtitlesWebvttCustomdataAPI communicates with '/encoding/configurations/subtitles/webvtt/{configuration_id}/customData' endpoints
type EncodingConfigurationsSubtitlesWebvttCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsSubtitlesWebvttCustomdataAPI constructor for EncodingConfigurationsSubtitlesWebvttCustomdataAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesWebvttCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesWebvttCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsSubtitlesWebvttCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesWebvttCustomdataAPIWithClient constructor for EncodingConfigurationsSubtitlesWebvttCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesWebvttCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesWebvttCustomdataAPI {
	a := &EncodingConfigurationsSubtitlesWebvttCustomdataAPI{apiClient: apiClient}
	return a
}

// Get WebVtt Subtitle Configuration Custom Data
func (api *EncodingConfigurationsSubtitlesWebvttCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/subtitles/webvtt/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
