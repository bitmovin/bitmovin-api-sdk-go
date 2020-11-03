package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmWidevineAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine' endpoints
type EncodingEncodingsMuxingsMp4DrmWidevineAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPI

}

// NewEncodingEncodingsMuxingsMp4DrmWidevineAPI constructor for EncodingEncodingsMuxingsMp4DrmWidevineAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmWidevineAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmWidevineAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmWidevineAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmWidevineAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmWidevineAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmWidevineAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmWidevineAPI {
    a := &EncodingEncodingsMuxingsMp4DrmWidevineAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Widevine DRM to an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmWidevineAPI) Create(encodingId string, muxingId string, widevineDrm model.WidevineDrm) (*model.WidevineDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.WidevineDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine", &widevineDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Widevine DRM from an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmWidevineAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Widevine DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmWidevineAPI) Get(encodingId string, muxingId string, drmId string) (*model.WidevineDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.WidevineDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Widevine DRMs of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmWidevineAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4DrmWidevineAPIListQueryParams)) (*pagination.WidevineDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsMp4DrmWidevineAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.WidevineDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsMp4DrmWidevineAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4DrmWidevineAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4DrmWidevineAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


