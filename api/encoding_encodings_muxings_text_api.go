package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsTextAPI communicates with '/encoding/encodings/{encoding_id}/muxings/text' endpoints
type EncodingEncodingsMuxingsTextAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsTextCustomdataAPI

}

// NewEncodingEncodingsMuxingsTextAPI constructor for EncodingEncodingsMuxingsTextAPI that takes options as argument
func NewEncodingEncodingsMuxingsTextAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTextAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsTextAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTextAPIWithClient constructor for EncodingEncodingsMuxingsTextAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTextAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTextAPI {
    a := &EncodingEncodingsMuxingsTextAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsTextCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Text muxing
func (api *EncodingEncodingsMuxingsTextAPI) Create(encodingId string, textMuxing model.TextMuxing) (*model.TextMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.TextMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/text", &textMuxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Text muxing
func (api *EncodingEncodingsMuxingsTextAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Text muxing details
func (api *EncodingEncodingsMuxingsTextAPI) Get(encodingId string, muxingId string) (*model.TextMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.TextMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Text muxings
func (api *EncodingEncodingsMuxingsTextAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsTextAPIListQueryParams)) (*pagination.TextMuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsTextAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.TextMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/text", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsTextAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsTextAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsTextAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


