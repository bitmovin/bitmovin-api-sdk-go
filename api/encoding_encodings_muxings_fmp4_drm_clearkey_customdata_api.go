package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI{apiClient: apiClient}
    return a
}

// Get ClearKey DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

