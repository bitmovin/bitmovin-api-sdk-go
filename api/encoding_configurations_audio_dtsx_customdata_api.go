package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioDtsxCustomdataAPI communicates with '/encoding/configurations/audio/dtsx/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioDtsxCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioDtsxCustomdataAPI constructor for EncodingConfigurationsAudioDtsxCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioDtsxCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDtsxCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsAudioDtsxCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDtsxCustomdataAPIWithClient constructor for EncodingConfigurationsAudioDtsxCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDtsxCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDtsxCustomdataAPI {
	a := &EncodingConfigurationsAudioDtsxCustomdataAPI{apiClient: apiClient}
	return a
}

// Get DTS:X Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioDtsxCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/audio/dtsx/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
