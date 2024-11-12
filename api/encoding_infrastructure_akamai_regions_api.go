package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureAkamaiRegionsAPI communicates with '/encoding/infrastructure/akamai/{infrastructure_id}/regions' endpoints
type EncodingInfrastructureAkamaiRegionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureAkamaiRegionsAPI constructor for EncodingInfrastructureAkamaiRegionsAPI that takes options as argument
func NewEncodingInfrastructureAkamaiRegionsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAkamaiRegionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureAkamaiRegionsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAkamaiRegionsAPIWithClient constructor for EncodingInfrastructureAkamaiRegionsAPI that takes an APIClient as argument
func NewEncodingInfrastructureAkamaiRegionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAkamaiRegionsAPI {
	a := &EncodingInfrastructureAkamaiRegionsAPI{apiClient: apiClient}
	return a
}

// Create Add Akamai account region settings
func (api *EncodingInfrastructureAkamaiRegionsAPI) Create(infrastructureId string, region model.AkamaiCloudRegion, akamaiAccountRegionSettings model.AkamaiAccountRegionSettings) (*model.AkamaiAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AkamaiAccountRegionSettings
	err := api.apiClient.Post("/encoding/infrastructure/akamai/{infrastructure_id}/regions/{region}", &akamaiAccountRegionSettings, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Akamai account region settings
func (api *EncodingInfrastructureAkamaiRegionsAPI) Delete(infrastructureId string, region model.AkamaiCloudRegion) (*model.AkamaiAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AkamaiAccountRegionSettings
	err := api.apiClient.Delete("/encoding/infrastructure/akamai/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Akamai account region settings details
func (api *EncodingInfrastructureAkamaiRegionsAPI) Get(infrastructureId string, region model.AkamaiCloudRegion) (*model.AkamaiAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.AkamaiAccountRegionSettings
	err := api.apiClient.Get("/encoding/infrastructure/akamai/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Akamai account region settings
func (api *EncodingInfrastructureAkamaiRegionsAPI) List(infrastructureId string, queryParams ...func(*EncodingInfrastructureAkamaiRegionsAPIListQueryParams)) (*pagination.AkamaiAccountRegionSettingssListPagination, error) {
	queryParameters := &EncodingInfrastructureAkamaiRegionsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AkamaiAccountRegionSettingssListPagination
	err := api.apiClient.Get("/encoding/infrastructure/akamai/{infrastructure_id}/regions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureAkamaiRegionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureAkamaiRegionsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureAkamaiRegionsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
