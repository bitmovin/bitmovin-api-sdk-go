package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI struct {
    apiClient *apiclient.APIClient

    // Contentprotection communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}/contentprotection' endpoints
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPI

}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI{apiClient: apiClient}
    a.Contentprotection = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionAPIWithClient(apiClient)

    return a
}

// Create Add DRM fMP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI) Create(manifestId string, periodId string, adaptationsetId string, dashFmp4DrmRepresentation model.DashFmp4DrmRepresentation) (*model.DashFmp4DrmRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.DashFmp4DrmRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm", &dashFmp4DrmRepresentation, &responseModel, reqParams)
    return &responseModel, err
}
// Delete DRM fMP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get DRM fMP4 Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashFmp4DrmRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.DashFmp4DrmRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all DRM fMP4 Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIListQueryParams)) (*pagination.DashFmp4DrmRepresentationsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DashFmp4DrmRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


