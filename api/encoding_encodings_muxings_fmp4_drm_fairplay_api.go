package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmFairplayAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay' endpoints
type EncodingEncodingsMuxingsFmp4DrmFairplayAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmFairplayAPI constructor for EncodingEncodingsMuxingsFmp4DrmFairplayAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmFairplayAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmFairplayAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmFairplayAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmFairplayAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmFairplayAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmFairplayAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmFairplayAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmFairplayAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add FairPlay DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmFairplayAPI) Create(encodingId string, muxingId string, fairPlayDrm model.FairPlayDrm) (*model.FairPlayDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.FairPlayDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay", &fairPlayDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete FairPlay DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmFairplayAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get FairPlay DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmFairplayAPI) Get(encodingId string, muxingId string, drmId string) (*model.FairPlayDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.FairPlayDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List FairPlay DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmFairplayAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmFairplayAPIListQueryParams)) (*pagination.FairPlayDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4DrmFairplayAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.FairPlayDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmFairplayAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmFairplayAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmFairplayAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


