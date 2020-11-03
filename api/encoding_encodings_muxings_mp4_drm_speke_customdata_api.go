package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI constructor for EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI {
    a := &EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI{apiClient: apiClient}
    return a
}

// Get SPEKE DRM Custom Data of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

