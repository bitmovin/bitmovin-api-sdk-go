package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsDrmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm' endpoints
type EncodingEncodingsMuxingsProgressiveTsDrmAPI struct {
    apiClient *apiclient.APIClient

    // Fairplay communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay' endpoints
    Fairplay *EncodingEncodingsMuxingsProgressiveTsDrmFairplayAPI
    // Aes communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/aes' endpoints
    Aes *EncodingEncodingsMuxingsProgressiveTsDrmAesAPI
    // Speke communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke' endpoints
    Speke *EncodingEncodingsMuxingsProgressiveTsDrmSpekeAPI

}

// NewEncodingEncodingsMuxingsProgressiveTsDrmAPI constructor for EncodingEncodingsMuxingsProgressiveTsDrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsDrmAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsDrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsDrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsDrmAPI {
    a := &EncodingEncodingsMuxingsProgressiveTsDrmAPI{apiClient: apiClient}
    a.Fairplay = NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayAPIWithClient(apiClient)
    a.Aes = NewEncodingEncodingsMuxingsProgressiveTsDrmAesAPIWithClient(apiClient)
    a.Speke = NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeAPIWithClient(apiClient)

    return a
}

// Get DRM Details of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmAPI) Get(encodingId string, muxingId string, drmId string) (*model.Drm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.Drm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all DRM configurations of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmAPI) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm", nil, &responseModel, reqParams)
    return &responseModel, err
}

