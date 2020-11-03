package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmSpekeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke' endpoints
type EncodingEncodingsMuxingsFmp4DrmSpekeAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmSpekeAPI constructor for EncodingEncodingsMuxingsFmp4DrmSpekeAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmSpekeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmSpekeAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmSpekeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmSpekeAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmSpekeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmSpekeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmSpekeAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmSpekeAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add SPEKE DRM key provider to fMP4
func (api *EncodingEncodingsMuxingsFmp4DrmSpekeAPI) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.SpekeDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete SPEKE DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmSpekeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get SPEKE DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmSpekeAPI) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.SpekeDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List SPEKE DRM of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmSpekeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmSpekeAPIListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4DrmSpekeAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SpekeDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmSpekeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmSpekeAPIListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmSpekeAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


