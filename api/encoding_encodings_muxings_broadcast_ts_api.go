package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsBroadcastTsAPI communicates with '/encoding/encodings/{encoding_id}/muxings/broadcast-ts' endpoints
type EncodingEncodingsMuxingsBroadcastTsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsBroadcastTsCustomdataAPI
    // Information communicates with '/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/information' endpoints
    Information *EncodingEncodingsMuxingsBroadcastTsInformationAPI

}

// NewEncodingEncodingsMuxingsBroadcastTsAPI constructor for EncodingEncodingsMuxingsBroadcastTsAPI that takes options as argument
func NewEncodingEncodingsMuxingsBroadcastTsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsBroadcastTsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsBroadcastTsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsBroadcastTsAPIWithClient constructor for EncodingEncodingsMuxingsBroadcastTsAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsBroadcastTsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsBroadcastTsAPI {
    a := &EncodingEncodingsMuxingsBroadcastTsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsBroadcastTsCustomdataAPIWithClient(apiClient)
    a.Information = NewEncodingEncodingsMuxingsBroadcastTsInformationAPIWithClient(apiClient)

    return a
}

// Create Add Broadcast TS muxing
func (api *EncodingEncodingsMuxingsBroadcastTsAPI) Create(encodingId string, broadcastTsMuxing model.BroadcastTsMuxing) (*model.BroadcastTsMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.BroadcastTsMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/broadcast-ts", &broadcastTsMuxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Broadcast TS muxing
func (api *EncodingEncodingsMuxingsBroadcastTsAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Broadcast TS muxing details
func (api *EncodingEncodingsMuxingsBroadcastTsAPI) Get(encodingId string, muxingId string) (*model.BroadcastTsMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BroadcastTsMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Broadcast TS muxings
func (api *EncodingEncodingsMuxingsBroadcastTsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsBroadcastTsAPIListQueryParams)) (*pagination.BroadcastTsMuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsBroadcastTsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.BroadcastTsMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsBroadcastTsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsBroadcastTsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsBroadcastTsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


