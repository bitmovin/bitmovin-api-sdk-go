package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsTsDrmSpekeAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke' endpoints
type EncodingEncodingsMuxingsTsDrmSpekeAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsTsDrmSpekeCustomdataAPI

}

// NewEncodingEncodingsMuxingsTsDrmSpekeAPI constructor for EncodingEncodingsMuxingsTsDrmSpekeAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmSpekeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmSpekeAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsTsDrmSpekeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmSpekeAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmSpekeAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmSpekeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmSpekeAPI {
    a := &EncodingEncodingsMuxingsTsDrmSpekeAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add SPEKE DRM key provider to a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmSpekeAPI) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.SpekeDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete SPEKE DRM from a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmSpekeAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get SPEKE DRM Details of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmSpekeAPI) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.SpekeDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List SPEKE DRM of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmSpekeAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsTsDrmSpekeAPIListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsTsDrmSpekeAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SpekeDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsTsDrmSpekeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsTsDrmSpekeAPIListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsTsDrmSpekeAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


