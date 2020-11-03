package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4API communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4' endpoints
type EncodingEncodingsMuxingsFmp4API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4CustomdataAPI
    // Information communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/information' endpoints
    Information *EncodingEncodingsMuxingsFmp4InformationAPI
    // Drm communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm' endpoints
    Drm *EncodingEncodingsMuxingsFmp4DrmAPI

}

// NewEncodingEncodingsMuxingsFmp4API constructor for EncodingEncodingsMuxingsFmp4API that takes options as argument
func NewEncodingEncodingsMuxingsFmp4API(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4APIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4APIWithClient constructor for EncodingEncodingsMuxingsFmp4API that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4APIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4API {
    a := &EncodingEncodingsMuxingsFmp4API{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4CustomdataAPIWithClient(apiClient)
    a.Information = NewEncodingEncodingsMuxingsFmp4InformationAPIWithClient(apiClient)
    a.Drm = NewEncodingEncodingsMuxingsFmp4DrmAPIWithClient(apiClient)

    return a
}

// Create Add fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4API) Create(encodingId string, fmp4Muxing model.Fmp4Muxing) (*model.Fmp4Muxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.Fmp4Muxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4", &fmp4Muxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4API) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get fMP4 muxing details
func (api *EncodingEncodingsMuxingsFmp4API) Get(encodingId string, muxingId string) (*model.Fmp4Muxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.Fmp4Muxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List fMP4 muxings
func (api *EncodingEncodingsMuxingsFmp4API) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4APIListQueryParams)) (*pagination.Fmp4MuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Fmp4MuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4APIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


