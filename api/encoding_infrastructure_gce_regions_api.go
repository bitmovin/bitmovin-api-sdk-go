package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureGceRegionsAPI communicates with '/encoding/infrastructure/gce/{infrastructure_id}/regions' endpoints
type EncodingInfrastructureGceRegionsAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureGceRegionsAPI constructor for EncodingInfrastructureGceRegionsAPI that takes options as argument
func NewEncodingInfrastructureGceRegionsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureGceRegionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInfrastructureGceRegionsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureGceRegionsAPIWithClient constructor for EncodingInfrastructureGceRegionsAPI that takes an APIClient as argument
func NewEncodingInfrastructureGceRegionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureGceRegionsAPI {
    a := &EncodingInfrastructureGceRegionsAPI{apiClient: apiClient}
    return a
}

// Create Add Google Cloud Region Setting
func (api *EncodingInfrastructureGceRegionsAPI) Create(infrastructureId string, region model.GoogleCloudRegion, gceAccountRegionSettings model.GceAccountRegionSettings) (*model.GceAccountRegionSettings, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel model.GceAccountRegionSettings
    err := api.apiClient.Post("/encoding/infrastructure/gce/{infrastructure_id}/regions/{region}", &gceAccountRegionSettings, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Google Cloud Region Settings
func (api *EncodingInfrastructureGceRegionsAPI) Delete(infrastructureId string, region model.GoogleCloudRegion) (*model.GceAccountRegionSettings, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel model.GceAccountRegionSettings
    err := api.apiClient.Delete("/encoding/infrastructure/gce/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Google Cloud Region Settings Details
func (api *EncodingInfrastructureGceRegionsAPI) Get(infrastructureId string, region model.GoogleCloudRegion) (*model.GceAccountRegionSettings, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel model.GceAccountRegionSettings
    err := api.apiClient.Get("/encoding/infrastructure/gce/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Google Cloud Region Settings
func (api *EncodingInfrastructureGceRegionsAPI) List(infrastructureId string, queryParams ...func(*EncodingInfrastructureGceRegionsAPIListQueryParams)) (*pagination.GceAccountRegionSettingssListPagination, error) {
    queryParameters := &EncodingInfrastructureGceRegionsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.GceAccountRegionSettingssListPagination
    err := api.apiClient.Get("/encoding/infrastructure/gce/{infrastructure_id}/regions", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInfrastructureGceRegionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureGceRegionsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureGceRegionsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


