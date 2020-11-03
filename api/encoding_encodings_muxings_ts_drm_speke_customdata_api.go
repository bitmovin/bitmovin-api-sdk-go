package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI constructor for EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI {
    a := &EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI{apiClient: apiClient}
    return a
}

// Get SPEKE DRM Custom Data of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

