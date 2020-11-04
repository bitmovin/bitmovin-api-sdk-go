package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureAzureRegionsAPI communicates with '/encoding/infrastructure/azure/{infrastructure_id}/regions' endpoints
type EncodingInfrastructureAzureRegionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureAzureRegionsAPI constructor for EncodingInfrastructureAzureRegionsAPI that takes options as argument
func NewEncodingInfrastructureAzureRegionsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAzureRegionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureAzureRegionsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAzureRegionsAPIWithClient constructor for EncodingInfrastructureAzureRegionsAPI that takes an APIClient as argument
func NewEncodingInfrastructureAzureRegionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAzureRegionsAPI {
	a := &EncodingInfrastructureAzureRegionsAPI{apiClient: apiClient}
	return a
}

// Create Add Azure Region Setting
func (api *EncodingInfrastructureAzureRegionsAPI) Create(infrastructureId string, region model.AzureCloudRegion, azureAccountRegionSettings model.AzureAccountRegionSettings) (*model.AzureAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AzureAccountRegionSettings
	err := api.apiClient.Post("/encoding/infrastructure/azure/{infrastructure_id}/regions/{region}", &azureAccountRegionSettings, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Azure Region Settings
func (api *EncodingInfrastructureAzureRegionsAPI) Delete(infrastructureId string, region model.AzureCloudRegion) (*model.AzureAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AzureAccountRegionSettings
	err := api.apiClient.Delete("/encoding/infrastructure/azure/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Azure Region Settings Details
func (api *EncodingInfrastructureAzureRegionsAPI) Get(infrastructureId string, region model.AzureCloudRegion) (*model.AzureAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AzureAccountRegionSettings
	err := api.apiClient.Get("/encoding/infrastructure/azure/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Azure Region Settings
func (api *EncodingInfrastructureAzureRegionsAPI) List(infrastructureId string, queryParams ...func(*EncodingInfrastructureAzureRegionsAPIListQueryParams)) (*pagination.AzureAccountRegionSettingssListPagination, error) {
	queryParameters := &EncodingInfrastructureAzureRegionsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AzureAccountRegionSettingssListPagination
	err := api.apiClient.Get("/encoding/infrastructure/azure/{infrastructure_id}/regions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureAzureRegionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureAzureRegionsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureAzureRegionsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
