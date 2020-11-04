package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoVp8CustomdataAPI communicates with '/encoding/configurations/video/vp8/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoVp8CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoVp8CustomdataAPI constructor for EncodingConfigurationsVideoVp8CustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoVp8CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoVp8CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoVp8CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoVp8CustomdataAPIWithClient constructor for EncodingConfigurationsVideoVp8CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoVp8CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoVp8CustomdataAPI {
	a := &EncodingConfigurationsVideoVp8CustomdataAPI{apiClient: apiClient}
	return a
}

// Get VP8 Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoVp8CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/video/vp8/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
