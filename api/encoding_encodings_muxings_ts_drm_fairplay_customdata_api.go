package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI constructor for EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsTsDrmFairplayCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmFairplayCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmFairplayCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI {
    a := &EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI{apiClient: apiClient}
    return a
}

// Get FairPlay DRM Custom Data of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmFairplayCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/fairplay/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

