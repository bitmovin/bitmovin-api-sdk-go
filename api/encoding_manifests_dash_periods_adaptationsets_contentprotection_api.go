package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI{apiClient: apiClient}
    return a
}

// Create Add Content Protection to AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI) Create(manifestId string, periodId string, adaptationsetId string, contentProtection model.ContentProtection) (*model.ContentProtection, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.ContentProtection
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection", &contentProtection, &responseModel, reqParams)
    return &responseModel, err
}
// Delete AdaptationSet Content Protection
func (api *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI) Delete(manifestId string, periodId string, adaptationsetId string, contentprotectionId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["contentprotection_id"] = contentprotectionId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get AdaptationSet Content Protection Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI) Get(manifestId string, periodId string, adaptationsetId string, contentprotectionId string) (*model.ContentProtection, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["contentprotection_id"] = contentprotectionId
    }

    var responseModel model.ContentProtection
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all AdaptationSet Content Protections
func (api *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIListQueryParams)) (*pagination.ContentProtectionsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ContentProtectionsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


