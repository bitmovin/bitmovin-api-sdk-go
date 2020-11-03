package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveMovAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-mov' endpoints
type EncodingEncodingsMuxingsProgressiveMovAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsProgressiveMovCustomdataAPI
    // Information communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/information' endpoints
    Information *EncodingEncodingsMuxingsProgressiveMovInformationAPI

}

// NewEncodingEncodingsMuxingsProgressiveMovAPI constructor for EncodingEncodingsMuxingsProgressiveMovAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveMovAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveMovAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveMovAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveMovAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveMovAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveMovAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveMovAPI {
    a := &EncodingEncodingsMuxingsProgressiveMovAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsProgressiveMovCustomdataAPIWithClient(apiClient)
    a.Information = NewEncodingEncodingsMuxingsProgressiveMovInformationAPIWithClient(apiClient)

    return a
}

// Create Add Progressive MOV muxing
func (api *EncodingEncodingsMuxingsProgressiveMovAPI) Create(encodingId string, progressiveMovMuxing model.ProgressiveMovMuxing) (*model.ProgressiveMovMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.ProgressiveMovMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-mov", &progressiveMovMuxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Progressive MOV muxing
func (api *EncodingEncodingsMuxingsProgressiveMovAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Progressive MOV muxing details
func (api *EncodingEncodingsMuxingsProgressiveMovAPI) Get(encodingId string, muxingId string) (*model.ProgressiveMovMuxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.ProgressiveMovMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Progressive MOV muxings
func (api *EncodingEncodingsMuxingsProgressiveMovAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveMovAPIListQueryParams)) (*pagination.ProgressiveMovMuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsProgressiveMovAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ProgressiveMovMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveMovAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveMovAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveMovAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


