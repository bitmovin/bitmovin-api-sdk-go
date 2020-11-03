package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI{apiClient: apiClient}
    return a
}

// Create Add Progressive WebM Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI) Create(manifestId string, periodId string, adaptationsetId string, dashProgressiveWebmRepresentation model.DashProgressiveWebmRepresentation) (*model.DashProgressiveWebmRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.DashProgressiveWebmRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm", &dashProgressiveWebmRepresentation, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Progressive WebM Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Progressive WebM Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashProgressiveWebmRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.DashProgressiveWebmRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all Progressive WebM Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIListQueryParams)) (*pagination.DashProgressiveWebmRepresentationsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DashProgressiveWebmRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


