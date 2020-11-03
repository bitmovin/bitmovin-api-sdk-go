package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4API communicates with '/encoding/encodings/{encoding_id}/muxings/mp4' endpoints
type EncodingEncodingsMuxingsMp4API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsMp4CustomdataAPI
    // Information communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/information' endpoints
    Information *EncodingEncodingsMuxingsMp4InformationAPI
    // Drm communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm' endpoints
    Drm *EncodingEncodingsMuxingsMp4DrmAPI

}

// NewEncodingEncodingsMuxingsMp4API constructor for EncodingEncodingsMuxingsMp4API that takes options as argument
func NewEncodingEncodingsMuxingsMp4API(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4APIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4APIWithClient constructor for EncodingEncodingsMuxingsMp4API that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4APIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4API {
    a := &EncodingEncodingsMuxingsMp4API{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsMp4CustomdataAPIWithClient(apiClient)
    a.Information = NewEncodingEncodingsMuxingsMp4InformationAPIWithClient(apiClient)
    a.Drm = NewEncodingEncodingsMuxingsMp4DrmAPIWithClient(apiClient)

    return a
}

// Create Add MP4 muxing
func (api *EncodingEncodingsMuxingsMp4API) Create(encodingId string, mp4Muxing model.Mp4Muxing) (*model.Mp4Muxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.Mp4Muxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4", &mp4Muxing, &responseModel, reqParams)
    return &responseModel, err
}
// Delete MP4 muxing
func (api *EncodingEncodingsMuxingsMp4API) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get MP4 muxing details
func (api *EncodingEncodingsMuxingsMp4API) Get(encodingId string, muxingId string) (*model.Mp4Muxing, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.Mp4Muxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List MP4 muxings
func (api *EncodingEncodingsMuxingsMp4API) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4APIListQueryParams)) (*pagination.Mp4MuxingsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsMp4APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Mp4MuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsMp4APIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


