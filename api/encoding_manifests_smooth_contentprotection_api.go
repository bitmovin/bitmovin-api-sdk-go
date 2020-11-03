package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsSmoothContentprotectionAPI communicates with '/encoding/manifests/smooth/{manifest_id}/contentprotection' endpoints
type EncodingManifestsSmoothContentprotectionAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsSmoothContentprotectionAPI constructor for EncodingManifestsSmoothContentprotectionAPI that takes options as argument
func NewEncodingManifestsSmoothContentprotectionAPI(options ...apiclient.APIClientOption) (*EncodingManifestsSmoothContentprotectionAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsSmoothContentprotectionAPIWithClient(apiClient), nil
}

// NewEncodingManifestsSmoothContentprotectionAPIWithClient constructor for EncodingManifestsSmoothContentprotectionAPI that takes an APIClient as argument
func NewEncodingManifestsSmoothContentprotectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsSmoothContentprotectionAPI {
    a := &EncodingManifestsSmoothContentprotectionAPI{apiClient: apiClient}
    return a
}

// Create Add Content Protection to Smooth Streaming
func (api *EncodingManifestsSmoothContentprotectionAPI) Create(manifestId string, smoothManifestContentProtection model.SmoothManifestContentProtection) (*model.SmoothManifestContentProtection, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.SmoothManifestContentProtection
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/contentprotection", &smoothManifestContentProtection, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Content Protection of Smooth Streaming
func (api *EncodingManifestsSmoothContentprotectionAPI) Delete(manifestId string, protectionId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["protection_id"] = protectionId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}/contentprotection/{protection_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Content Protection of Smooth Streaming Representation Details
func (api *EncodingManifestsSmoothContentprotectionAPI) Get(manifestId string, protectionId string) (*model.SmoothManifestContentProtection, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["protection_id"] = protectionId
    }

    var responseModel model.SmoothManifestContentProtection
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/contentprotection/{protection_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Content Protection of Smooth Streaming
func (api *EncodingManifestsSmoothContentprotectionAPI) List(manifestId string, queryParams ...func(*EncodingManifestsSmoothContentprotectionAPIListQueryParams)) (*pagination.SmoothManifestContentProtectionsListPagination, error) {
    queryParameters := &EncodingManifestsSmoothContentprotectionAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SmoothManifestContentProtectionsListPagination
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/contentprotection", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsSmoothContentprotectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsSmoothContentprotectionAPIListQueryParams struct {
    Offset int64 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsSmoothContentprotectionAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


