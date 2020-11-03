package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI{apiClient: apiClient}
    return a
}

// Create Add VTT Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI) Create(manifestId string, periodId string, adaptationsetId string, dashVttRepresentation model.DashVttRepresentation) (*model.DashVttRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.DashVttRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt", &dashVttRepresentation, &responseModel, reqParams)
    return &responseModel, err
}
// Delete VTT Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get VTT Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashVttRepresentation, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel model.DashVttRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt/{representation_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all VTT Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIListQueryParams)) (*pagination.DashVttRepresentationsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DashVttRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


