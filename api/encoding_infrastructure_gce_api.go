package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureGceAPI communicates with '/encoding/infrastructure/gce' endpoints
type EncodingInfrastructureGceAPI struct {
	apiClient *apiclient.APIClient

	// Regions communicates with '/encoding/infrastructure/gce/{infrastructure_id}/regions' endpoints
	Regions *EncodingInfrastructureGceRegionsAPI
}

// NewEncodingInfrastructureGceAPI constructor for EncodingInfrastructureGceAPI that takes options as argument
func NewEncodingInfrastructureGceAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureGceAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureGceAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureGceAPIWithClient constructor for EncodingInfrastructureGceAPI that takes an APIClient as argument
func NewEncodingInfrastructureGceAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureGceAPI {
	a := &EncodingInfrastructureGceAPI{apiClient: apiClient}
	a.Regions = NewEncodingInfrastructureGceRegionsAPIWithClient(apiClient)

	return a
}

// Create Add GCE Account
func (api *EncodingInfrastructureGceAPI) Create(gceAccount model.GceAccount) (*model.GceAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.GceAccount
	err := api.apiClient.Post("/encoding/infrastructure/gce", &gceAccount, &responseModel, reqParams)
	return &responseModel, err
}

// Delete GCE Account
func (api *EncodingInfrastructureGceAPI) Delete(infrastructureId string) (*model.GceAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.GceAccount
	err := api.apiClient.Delete("/encoding/infrastructure/gce/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get GCE Account Details
func (api *EncodingInfrastructureGceAPI) Get(infrastructureId string) (*model.GceAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.GceAccount
	err := api.apiClient.Get("/encoding/infrastructure/gce/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List GCE Accounts
func (api *EncodingInfrastructureGceAPI) List(queryParams ...func(*EncodingInfrastructureGceAPIListQueryParams)) (*pagination.GceAccountsListPagination, error) {
	queryParameters := &EncodingInfrastructureGceAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.GceAccountsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/gce", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureGceAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureGceAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureGceAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
