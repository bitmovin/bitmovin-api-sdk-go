package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready' endpoints
type EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmPlayreadyAPI constructor for EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmPlayreadyAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmPlayreadyAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmPlayreadyAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmPlayreadyAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add PlayReady DRM to an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI) Create(encodingId string, muxingId string, playReadyDrm model.PlayReadyDrm) (*model.PlayReadyDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.PlayReadyDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready", &playReadyDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete PlayReady DRM from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get PlayReady DRM Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI) Get(encodingId string, muxingId string, drmId string) (*model.PlayReadyDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.PlayReadyDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List PlayReady DRMs of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmPlayreadyAPIListQueryParams)) (*pagination.PlayReadyDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4DrmPlayreadyAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.PlayReadyDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmPlayreadyAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmPlayreadyAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmPlayreadyAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


