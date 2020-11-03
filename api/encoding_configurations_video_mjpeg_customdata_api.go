package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoMjpegCustomdataAPI communicates with '/encoding/configurations/video/mjpeg/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoMjpegCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoMjpegCustomdataAPI constructor for EncodingConfigurationsVideoMjpegCustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoMjpegCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoMjpegCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoMjpegCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoMjpegCustomdataAPIWithClient constructor for EncodingConfigurationsVideoMjpegCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoMjpegCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoMjpegCustomdataAPI {
    a := &EncodingConfigurationsVideoMjpegCustomdataAPI{apiClient: apiClient}
    return a
}

// Get MJPEG Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoMjpegCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/mjpeg/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

