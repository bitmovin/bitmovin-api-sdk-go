package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsBifsAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs' endpoints
type EncodingEncodingsStreamsBifsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}/customData' endpoints
    Customdata *EncodingEncodingsStreamsBifsCustomdataAPI

}

// NewEncodingEncodingsStreamsBifsAPI constructor for EncodingEncodingsStreamsBifsAPI that takes options as argument
func NewEncodingEncodingsStreamsBifsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsBifsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsBifsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsBifsAPIWithClient constructor for EncodingEncodingsStreamsBifsAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsBifsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsBifsAPI {
    a := &EncodingEncodingsStreamsBifsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsStreamsBifsCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add a Roku Bif file
func (api *EncodingEncodingsStreamsBifsAPI) Create(encodingId string, streamId string, bif model.Bif) (*model.Bif, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.Bif
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs", &bif, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Bif
func (api *EncodingEncodingsStreamsBifsAPI) Delete(encodingId string, streamId string, bifId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["bif_id"] = bifId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Bif Details
func (api *EncodingEncodingsStreamsBifsAPI) Get(encodingId string, streamId string, bifId string) (*model.Bif, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["bif_id"] = bifId
    }

    var responseModel model.Bif
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Bifs
func (api *EncodingEncodingsStreamsBifsAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsBifsAPIListQueryParams)) (*pagination.BifsListPagination, error) {
    queryParameters := &EncodingEncodingsStreamsBifsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.BifsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsStreamsBifsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsBifsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsBifsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


