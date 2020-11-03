package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm' endpoints
type EncodingEncodingsMuxingsFmp4DrmAPI struct {
    apiClient *apiclient.APIClient

    // Widevine communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine' endpoints
    Widevine *EncodingEncodingsMuxingsFmp4DrmWidevineAPI
    // Playready communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready' endpoints
    Playready *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI
    // Primetime communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/primetime' endpoints
    Primetime *EncodingEncodingsMuxingsFmp4DrmPrimetimeAPI
    // Fairplay communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay' endpoints
    Fairplay *EncodingEncodingsMuxingsFmp4DrmFairplayAPI
    // Marlin communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin' endpoints
    Marlin *EncodingEncodingsMuxingsFmp4DrmMarlinAPI
    // Clearkey communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey' endpoints
    Clearkey *EncodingEncodingsMuxingsFmp4DrmClearkeyAPI
    // Cenc communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc' endpoints
    Cenc *EncodingEncodingsMuxingsFmp4DrmCencAPI
    // Aes communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes' endpoints
    Aes *EncodingEncodingsMuxingsFmp4DrmAesAPI
    // Speke communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke' endpoints
    Speke *EncodingEncodingsMuxingsFmp4DrmSpekeAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmAPI constructor for EncodingEncodingsMuxingsFmp4DrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmAPI{apiClient: apiClient}
    a.Widevine = NewEncodingEncodingsMuxingsFmp4DrmWidevineAPIWithClient(apiClient)
    a.Playready = NewEncodingEncodingsMuxingsFmp4DrmPlayreadyAPIWithClient(apiClient)
    a.Primetime = NewEncodingEncodingsMuxingsFmp4DrmPrimetimeAPIWithClient(apiClient)
    a.Fairplay = NewEncodingEncodingsMuxingsFmp4DrmFairplayAPIWithClient(apiClient)
    a.Marlin = NewEncodingEncodingsMuxingsFmp4DrmMarlinAPIWithClient(apiClient)
    a.Clearkey = NewEncodingEncodingsMuxingsFmp4DrmClearkeyAPIWithClient(apiClient)
    a.Cenc = NewEncodingEncodingsMuxingsFmp4DrmCencAPIWithClient(apiClient)
    a.Aes = NewEncodingEncodingsMuxingsFmp4DrmAesAPIWithClient(apiClient)
    a.Speke = NewEncodingEncodingsMuxingsFmp4DrmSpekeAPIWithClient(apiClient)

    return a
}

// Get DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmAPI) Get(encodingId string, muxingId string, drmId string) (*model.Drm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.Drm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmAPI) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm", nil, &responseModel, reqParams)
    return &responseModel, err
}

