package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI {
    a := &EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI{apiClient: apiClient}
    return a
}

// Get FairPlay DRM Custom Data of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

