package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI {
    a := &EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI{apiClient: apiClient}
    return a
}

// Get CENC DRM Custom Data of a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmCencCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/cenc/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

