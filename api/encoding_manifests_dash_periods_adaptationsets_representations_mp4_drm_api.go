package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI{apiClient: apiClient}
    return a
}

// Create Add DRM MP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI) Create(manifestId string, periodId string, adaptationsetId string, dashMp4DrmRepresentation model.DashMp4DrmRepresentation) (*model.DashMp4DrmRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.DashMp4DrmRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm", &dashMp4DrmRepresentation, &responseModel, reqParams)
    return &responseModel, err
}
// Delete DRM MP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get DRM MP4 Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashMp4DrmRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.DashMp4DrmRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all DRM MP4 Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIListQueryParams)) (*pagination.DashMp4DrmRepresentationsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DashMp4DrmRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


