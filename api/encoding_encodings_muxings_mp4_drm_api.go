package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm' endpoints
type EncodingEncodingsMuxingsMp4DrmAPI struct {
    apiClient *apiclient.APIClient

    // Playready communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready' endpoints
    Playready *EncodingEncodingsMuxingsMp4DrmPlayreadyAPI
    // Clearkey communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey' endpoints
    Clearkey *EncodingEncodingsMuxingsMp4DrmClearkeyAPI
    // Widevine communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine' endpoints
    Widevine *EncodingEncodingsMuxingsMp4DrmWidevineAPI
    // Marlin communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin' endpoints
    Marlin *EncodingEncodingsMuxingsMp4DrmMarlinAPI
    // Cenc communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc' endpoints
    Cenc *EncodingEncodingsMuxingsMp4DrmCencAPI
    // Speke communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke' endpoints
    Speke *EncodingEncodingsMuxingsMp4DrmSpekeAPI

}

// NewEncodingEncodingsMuxingsMp4DrmAPI constructor for EncodingEncodingsMuxingsMp4DrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmAPI {
    a := &EncodingEncodingsMuxingsMp4DrmAPI{apiClient: apiClient}
    a.Playready = NewEncodingEncodingsMuxingsMp4DrmPlayreadyAPIWithClient(apiClient)
    a.Clearkey = NewEncodingEncodingsMuxingsMp4DrmClearkeyAPIWithClient(apiClient)
    a.Widevine = NewEncodingEncodingsMuxingsMp4DrmWidevineAPIWithClient(apiClient)
    a.Marlin = NewEncodingEncodingsMuxingsMp4DrmMarlinAPIWithClient(apiClient)
    a.Cenc = NewEncodingEncodingsMuxingsMp4DrmCencAPIWithClient(apiClient)
    a.Speke = NewEncodingEncodingsMuxingsMp4DrmSpekeAPIWithClient(apiClient)

    return a
}

// Get DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmAPI) Get(encodingId string, muxingId string, drmId string) (*model.Drm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.Drm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all DRM configurations of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmAPI) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm", nil, &responseModel, reqParams)
    return &responseModel, err
}

