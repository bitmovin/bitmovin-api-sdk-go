package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmClearkeyAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey' endpoints
type EncodingEncodingsMuxingsFmp4DrmClearkeyAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmClearkeyAPI constructor for EncodingEncodingsMuxingsFmp4DrmClearkeyAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmClearkeyAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmClearkeyAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmClearkeyAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmClearkeyAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmClearkeyAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmClearkeyAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmClearkeyAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmClearkeyAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add ClearKey DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyAPI) Create(encodingId string, muxingId string, clearKeyDrm model.ClearKeyDrm) (*model.ClearKeyDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.ClearKeyDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey", &clearKeyDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete ClearKey DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get ClearKey DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyAPI) Get(encodingId string, muxingId string, drmId string) (*model.ClearKeyDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.ClearKeyDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List ClearKey DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmClearkeyAPIListQueryParams)) (*pagination.ClearKeyDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4DrmClearkeyAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ClearKeyDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmClearkeyAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmClearkeyAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmClearkeyAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


