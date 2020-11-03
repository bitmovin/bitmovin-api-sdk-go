package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI constructor for EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI {
    a := &EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Widevine DRM Custom Data of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

