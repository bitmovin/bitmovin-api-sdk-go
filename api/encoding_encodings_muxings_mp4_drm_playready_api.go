package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsMp4DrmPlayreadyAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready' endpoints
type EncodingEncodingsMuxingsMp4DrmPlayreadyAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI

}

// NewEncodingEncodingsMuxingsMp4DrmPlayreadyAPI constructor for EncodingEncodingsMuxingsMp4DrmPlayreadyAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmPlayreadyAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmPlayreadyAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsMp4DrmPlayreadyAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmPlayreadyAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmPlayreadyAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmPlayreadyAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmPlayreadyAPI {
    a := &EncodingEncodingsMuxingsMp4DrmPlayreadyAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add PlayReady DRM to an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyAPI) Create(encodingId string, muxingId string, playReadyDrm model.PlayReadyDrm) (*model.PlayReadyDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.PlayReadyDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready", &playReadyDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete PlayReady DRM from an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get PlayReady DRM Details of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyAPI) Get(encodingId string, muxingId string, drmId string) (*model.PlayReadyDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.PlayReadyDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List PlayReady DRMs of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsMp4DrmPlayreadyAPIListQueryParams)) (*pagination.PlayReadyDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsMp4DrmPlayreadyAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.PlayReadyDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsMp4DrmPlayreadyAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsMp4DrmPlayreadyAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsMp4DrmPlayreadyAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


