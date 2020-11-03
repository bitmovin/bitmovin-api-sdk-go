package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API struct {
    apiClient *apiclient.APIClient

    // Drm communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm' endpoints
    Drm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPI
    // Contentprotection communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection' endpoints
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPI

}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API {
    a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API{apiClient: apiClient}
    a.Drm = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmAPIWithClient(apiClient)
    a.Contentprotection = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionAPIWithClient(apiClient)

    return a
}

// Create Add fMP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API) Create(manifestId string, periodId string, adaptationsetId string, dashFmp4Representation model.DashFmp4Representation) (*model.DashFmp4Representation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.DashFmp4Representation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4", &dashFmp4Representation, &responseModel, reqParams)
    return &responseModel, err
}
// Delete fMP4 Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get fMP4 Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashFmp4Representation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.DashFmp4Representation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all fMP4 Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIListQueryParams)) (*pagination.DashFmp4RepresentationsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DashFmp4RepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


