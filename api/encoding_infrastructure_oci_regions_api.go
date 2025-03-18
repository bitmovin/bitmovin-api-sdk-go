package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureOciRegionsAPI communicates with '/encoding/infrastructure/oci/{infrastructure_id}/regions' endpoints
type EncodingInfrastructureOciRegionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureOciRegionsAPI constructor for EncodingInfrastructureOciRegionsAPI that takes options as argument
func NewEncodingInfrastructureOciRegionsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureOciRegionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureOciRegionsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureOciRegionsAPIWithClient constructor for EncodingInfrastructureOciRegionsAPI that takes an APIClient as argument
func NewEncodingInfrastructureOciRegionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureOciRegionsAPI {
	a := &EncodingInfrastructureOciRegionsAPI{apiClient: apiClient}
	return a
}

// Create Add OCI account region settings
func (api *EncodingInfrastructureOciRegionsAPI) Create(infrastructureId string, region model.OciCloudRegion, ociAccountRegionSettings model.OciAccountRegionSettings) (*model.OciAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.OciAccountRegionSettings
	err := api.apiClient.Post("/encoding/infrastructure/oci/{infrastructure_id}/regions/{region}", &ociAccountRegionSettings, &responseModel, reqParams)
	return &responseModel, err
}

// Delete OCI account region settings
func (api *EncodingInfrastructureOciRegionsAPI) Delete(infrastructureId string, region model.OciCloudRegion) (*model.OciAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.OciAccountRegionSettings
	err := api.apiClient.Delete("/encoding/infrastructure/oci/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get OCI account region settings details
func (api *EncodingInfrastructureOciRegionsAPI) Get(infrastructureId string, region model.OciCloudRegion) (*model.OciAccountRegionSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.PathParams["region"] = string(region)
	}

	var responseModel model.OciAccountRegionSettings
	err := api.apiClient.Get("/encoding/infrastructure/oci/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List OCI account region settings
func (api *EncodingInfrastructureOciRegionsAPI) List(infrastructureId string, queryParams ...func(*EncodingInfrastructureOciRegionsAPIListQueryParams)) (*pagination.OciAccountRegionSettingssListPagination, error) {
	queryParameters := &EncodingInfrastructureOciRegionsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.OciAccountRegionSettingssListPagination
	err := api.apiClient.Get("/encoding/infrastructure/oci/{infrastructure_id}/regions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureOciRegionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureOciRegionsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureOciRegionsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
