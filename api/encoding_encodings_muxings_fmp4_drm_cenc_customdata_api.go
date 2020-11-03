package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmCencCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmCencCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmCencCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI{apiClient: apiClient}
    return a
}

// Get CENC DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

