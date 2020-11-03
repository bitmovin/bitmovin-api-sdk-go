package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsFmp4DrmAesAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes' endpoints
type EncodingEncodingsMuxingsFmp4DrmAesAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI

}

// NewEncodingEncodingsMuxingsFmp4DrmAesAPI constructor for EncodingEncodingsMuxingsFmp4DrmAesAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmAesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmAesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsFmp4DrmAesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmAesAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmAesAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmAesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmAesAPI {
    a := &EncodingEncodingsMuxingsFmp4DrmAesAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add AES encryption configuration to fMP4
func (api *EncodingEncodingsMuxingsFmp4DrmAesAPI) Create(encodingId string, muxingId string, aesEncryptionDrm model.AesEncryptionDrm) (*model.AesEncryptionDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.AesEncryptionDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes", &aesEncryptionDrm, &responseModel, reqParams)
    return &responseModel, err
}
// Delete AES encryption configuration from an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmAesAPI) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get AES encryption Details of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmAesAPI) Get(encodingId string, muxingId string, drmId string) (*model.AesEncryptionDrm, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel model.AesEncryptionDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List AES encryption configurations of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmAesAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsFmp4DrmAesAPIListQueryParams)) (*pagination.AesEncryptionDrmsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsFmp4DrmAesAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AesEncryptionDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsFmp4DrmAesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsFmp4DrmAesAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsFmp4DrmAesAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


