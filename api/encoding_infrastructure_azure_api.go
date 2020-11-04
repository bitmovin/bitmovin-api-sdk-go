package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureAzureAPI communicates with '/encoding/infrastructure/azure' endpoints
type EncodingInfrastructureAzureAPI struct {
	apiClient *apiclient.APIClient

	// Regions communicates with '/encoding/infrastructure/azure/{infrastructure_id}/regions' endpoints
	Regions *EncodingInfrastructureAzureRegionsAPI
}

// NewEncodingInfrastructureAzureAPI constructor for EncodingInfrastructureAzureAPI that takes options as argument
func NewEncodingInfrastructureAzureAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAzureAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureAzureAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAzureAPIWithClient constructor for EncodingInfrastructureAzureAPI that takes an APIClient as argument
func NewEncodingInfrastructureAzureAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAzureAPI {
	a := &EncodingInfrastructureAzureAPI{apiClient: apiClient}
	a.Regions = NewEncodingInfrastructureAzureRegionsAPIWithClient(apiClient)

	return a
}

// Create Add Azure Account
func (api *EncodingInfrastructureAzureAPI) Create(azureAccount model.AzureAccount) (*model.AzureAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AzureAccount
	err := api.apiClient.Post("/encoding/infrastructure/azure", &azureAccount, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Azure Account
func (api *EncodingInfrastructureAzureAPI) Delete(infrastructureId string) (*model.AzureAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.AzureAccount
	err := api.apiClient.Delete("/encoding/infrastructure/azure/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Azure Account Details
func (api *EncodingInfrastructureAzureAPI) Get(infrastructureId string) (*model.AzureAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.AzureAccount
	err := api.apiClient.Get("/encoding/infrastructure/azure/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Azure Accounts
func (api *EncodingInfrastructureAzureAPI) List(queryParams ...func(*EncodingInfrastructureAzureAPIListQueryParams)) (*pagination.AzureAccountsListPagination, error) {
	queryParameters := &EncodingInfrastructureAzureAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AzureAccountsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/azure", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureAzureAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureAzureAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureAzureAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
