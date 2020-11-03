package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsTsAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts' endpoints
type EncodingEncodingsMuxingsTsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsTsCustomdataAPI
    // Drm communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm' endpoints
    Drm *EncodingEncodingsMuxingsTsDrmAPI

}

// NewEncodingEncodingsMuxingsTsAPI constructor for EncodingEncodingsMuxingsTsAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsTsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsAPIWithClient constructor for EncodingEncodingsMuxingsTsAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsAPI {
    a := &EncodingEncodingsMuxingsTsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsTsCustomdataAPIWithClient(apiClient)
    a.Drm = NewEncodingEncodingsMuxingsTsDrmAPIWithClient(apiClient)

    return a
}

// Create Add TS muxing
func (api *EncodingEncodingsMuxingsTsAPI) Create(encodingId string, tsMuxing model.TsMuxing) (*model.TsMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.TsMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts", &tsMuxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete TS muxing
func (api *EncodingEncodingsMuxingsTsAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get TS muxing details
func (api *EncodingEncodingsMuxingsTsAPI) Get(encodingId string, muxingId string) (*model.TsMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.TsMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List TS muxings
func (api *EncodingEncodingsMuxingsTsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsTsAPIListQueryParams)) (*pagination.TsMuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsTsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.TsMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsTsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsTsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsTsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


