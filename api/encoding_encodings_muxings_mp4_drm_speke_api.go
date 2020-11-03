package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmSpekeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke' endpoints
type EncodingEncodingsMuxingsMp4DrmSpekeAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPI

}

// NewEncodingEncodingsMuxingsMp4DrmSpekeAPI constructor for EncodingEncodingsMuxingsMp4DrmSpekeAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmSpekeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmSpekeAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmSpekeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmSpekeAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmSpekeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmSpekeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmSpekeAPI {
    a := &EncodingEncodingsMuxingsMp4DrmSpekeAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add SPEKE DRM key provider to an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmSpekeAPI) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.SpekeDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete SPEKE DRM from an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmSpekeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get SPEKE DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmSpekeAPI) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.SpekeDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List SPEKE DRM of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmSpekeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4DrmSpekeAPIListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsMp4DrmSpekeAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SpekeDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsMp4DrmSpekeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4DrmSpekeAPIListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4DrmSpekeAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


