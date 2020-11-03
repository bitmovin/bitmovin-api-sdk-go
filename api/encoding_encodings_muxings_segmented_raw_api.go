package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsSegmentedRawAPI communicates with '/encoding/encodings/{encoding_id}/muxings/segmented-raw' endpoints
type EncodingEncodingsMuxingsSegmentedRawAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsSegmentedRawCustomdataAPI

}

// NewEncodingEncodingsMuxingsSegmentedRawAPI constructor for EncodingEncodingsMuxingsSegmentedRawAPI that takes options as argument
func NewEncodingEncodingsMuxingsSegmentedRawAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsSegmentedRawAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsSegmentedRawAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsSegmentedRawAPIWithClient constructor for EncodingEncodingsMuxingsSegmentedRawAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsSegmentedRawAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsSegmentedRawAPI {
    a := &EncodingEncodingsMuxingsSegmentedRawAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsSegmentedRawCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Segmented RAW muxing
func (api *EncodingEncodingsMuxingsSegmentedRawAPI) Create(encodingId string, segmentedRawMuxing model.SegmentedRawMuxing) (*model.SegmentedRawMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.SegmentedRawMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/segmented-raw", &segmentedRawMuxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Segmented RAW muxing
func (api *EncodingEncodingsMuxingsSegmentedRawAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Segmented RAW muxing details
func (api *EncodingEncodingsMuxingsSegmentedRawAPI) Get(encodingId string, muxingId string) (*model.SegmentedRawMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.SegmentedRawMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Segmented RAW muxings
func (api *EncodingEncodingsMuxingsSegmentedRawAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsSegmentedRawAPIListQueryParams)) (*pagination.SegmentedRawMuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsSegmentedRawAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SegmentedRawMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsSegmentedRawAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsSegmentedRawAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsSegmentedRawAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


