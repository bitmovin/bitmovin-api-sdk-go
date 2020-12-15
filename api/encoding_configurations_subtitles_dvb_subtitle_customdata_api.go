package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI communicates with '/encoding/configurations/subtitles/dvb-subtitle/{configuration_id}/customData' endpoints
type EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI constructor for EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPIWithClient constructor for EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI {
	a := &EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI{apiClient: apiClient}
	return a
}

// Get DVB-SUB subtitle configuration custom data
func (api *EncodingConfigurationsSubtitlesDvbSubtitleCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/subtitles/dvb-subtitle/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
