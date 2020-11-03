package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI constructor for EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI {
    a := &EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Marlin DRM Custom Data of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmMarlinCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

