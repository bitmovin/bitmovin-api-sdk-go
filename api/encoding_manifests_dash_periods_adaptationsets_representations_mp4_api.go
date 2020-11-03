package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API struct {
    apiClient *apiclient.APIClient

    // Drm communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm' endpoints
    Drm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPI

}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API {
    a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API{apiClient: apiClient}
    a.Drm = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmAPIWithClient(apiClient)

    return a
}

// Create Add MP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API) Create(manifestId string, periodId string, adaptationsetId string, dashMp4Representation model.DashMp4Representation) (*model.DashMp4Representation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.DashMp4Representation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4", &dashMp4Representation, &responseModel, reqParams)
    return &responseModel, err
}
// Delete MP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get MP4 Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashMp4Representation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.DashMp4Representation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all MP4 Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIListQueryParams)) (*pagination.DashMp4RepresentationsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DashMp4RepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


