package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI{apiClient: apiClient}
    return a
}

// Create Add Subtitle AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI) Create(manifestId string, periodId string, subtitleAdaptationSet model.SubtitleAdaptationSet) (*model.SubtitleAdaptationSet, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }

    var responseModel model.SubtitleAdaptationSet
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle", &subtitleAdaptationSet, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Subtitle AdaptationSet
func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI) Delete(manifestId string, periodId string, adaptationsetId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle/{adaptationset_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Subtitle AdaptationSet Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI) Get(manifestId string, periodId string, adaptationsetId string) (*model.SubtitleAdaptationSet, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel model.SubtitleAdaptationSet
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle/{adaptationset_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all Subtitle AdaptationSets
func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI) List(manifestId string, periodId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIListQueryParams)) (*pagination.SubtitleAdaptationSetsListPagination, error) {
    queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SubtitleAdaptationSetsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


