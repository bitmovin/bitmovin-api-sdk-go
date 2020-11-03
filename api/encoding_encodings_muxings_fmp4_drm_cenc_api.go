package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmCencAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc' endpoints
type EncodingEncodingsMuxingsFmp4DrmCencAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4DrmCencCustomdataAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmCencAPI constructor for EncodingEncodingsMuxingsFmp4DrmCencAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmCencAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmCencAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmCencAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmCencAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmCencAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmCencAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmCencAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmCencAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmCencCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add CENC DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmCencAPI) Create(encodingId string, muxingId string, cencDrm model.CencDrm) (*model.CencDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.CencDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc", &cencDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete CENC DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmCencAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get CENC DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmCencAPI) Get(encodingId string, muxingId string, drmId string) (*model.CencDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.CencDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List CENC DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmCencAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmCencAPIListQueryParams)) (*pagination.CencDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4DrmCencAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.CencDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/cenc", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmCencAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmCencAPIListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmCencAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


