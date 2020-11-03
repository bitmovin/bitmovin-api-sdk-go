package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI{apiClient: apiClient}
    return a
}

// Get PrimeTime DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPrimetimeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime/{drm_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

