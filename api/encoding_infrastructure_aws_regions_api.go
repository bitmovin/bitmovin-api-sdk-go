package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureAwsRegionsAPI communicates with '/encoding/infrastructure/aws/{infrastructure_id}/regions' endpoints
type EncodingInfrastructureAwsRegionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureAwsRegionsAPI constructor for EncodingInfrastructureAwsRegionsAPI that takes options as argument
func NewEncodingInfrastructureAwsRegionsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAwsRegionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureAwsRegionsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAwsRegionsAPIWithClient constructor for EncodingInfrastructureAwsRegionsAPI that takes an APIClient as argument
func NewEncodingInfrastructureAwsRegionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAwsRegionsAPI {
	a := &EncodingInfrastructureAwsRegionsAPI{apiClient: apiClient}
	return a
}

// Create Add AWS Region Setting
func (api *EncodingInfrastructureAwsRegionsAPI) Create(infrastructureId string, region model.AwsCloudRegion, awsAccountRegionSettings model.AwsAccountRegionSettings) (*model.AwsAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AwsAccountRegionSettings
	err := api.apiClient.Post("/encoding/infrastructure/aws/{infrastructure_id}/regions/{region}", &awsAccountRegionSettings, &responseModel, reqParams)
	return &responseModel, err
}

// Delete AWS Region Settings
func (api *EncodingInfrastructureAwsRegionsAPI) Delete(infrastructureId string, region model.AwsCloudRegion) (*model.AwsAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AwsAccountRegionSettings
	err := api.apiClient.Delete("/encoding/infrastructure/aws/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get AWS Region Settings Details
func (api *EncodingInfrastructureAwsRegionsAPI) Get(infrastructureId string, region model.AwsCloudRegion) (*model.AwsAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AwsAccountRegionSettings
	err := api.apiClient.Get("/encoding/infrastructure/aws/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List AWS Region Settings
func (api *EncodingInfrastructureAwsRegionsAPI) List(infrastructureId string, queryParams ...func(*EncodingInfrastructureAwsRegionsAPIListQueryParams)) (*pagination.AwsAccountRegionSettingssListPagination, error) {
	queryParameters := &EncodingInfrastructureAwsRegionsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AwsAccountRegionSettingssListPagination
	err := api.apiClient.Get("/encoding/infrastructure/aws/{infrastructure_id}/regions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureAwsRegionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureAwsRegionsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureAwsRegionsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
