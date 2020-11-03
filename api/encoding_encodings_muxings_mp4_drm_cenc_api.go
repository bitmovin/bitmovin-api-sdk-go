package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmCencAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc' endpoints
type EncodingEncodingsMuxingsMp4DrmCencAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI

}

// NewEncodingEncodingsMuxingsMp4DrmCencAPI constructor for EncodingEncodingsMuxingsMp4DrmCencAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmCencAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmCencAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmCencAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmCencAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmCencAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmCencAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmCencAPI {
    a := &EncodingEncodingsMuxingsMp4DrmCencAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsMp4DrmCencCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add CENC DRM to an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmCencAPI) Create(encodingId string, muxingId string, cencDrm model.CencDrm) (*model.CencDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CencDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc", &cencDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete CENC DRM from an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmCencAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get CENC DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmCencAPI) Get(encodingId string, muxingId string, drmId string) (*model.CencDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CencDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List CENC DRMs of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmCencAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4DrmCencAPIListQueryParams)) (*pagination.CencDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsMp4DrmCencAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.CencDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsMp4DrmCencAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4DrmCencAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4DrmCencAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


