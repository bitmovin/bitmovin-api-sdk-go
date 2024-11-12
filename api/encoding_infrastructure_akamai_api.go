package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureAkamaiAPI communicates with '/encoding/infrastructure/akamai' endpoints
type EncodingInfrastructureAkamaiAPI struct {
	apiClient *apiclient.APIClient

	// Regions communicates with '/encoding/infrastructure/akamai/{infrastructure_id}/regions' endpoints
	Regions *EncodingInfrastructureAkamaiRegionsAPI
}

// NewEncodingInfrastructureAkamaiAPI constructor for EncodingInfrastructureAkamaiAPI that takes options as argument
func NewEncodingInfrastructureAkamaiAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAkamaiAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureAkamaiAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAkamaiAPIWithClient constructor for EncodingInfrastructureAkamaiAPI that takes an APIClient as argument
func NewEncodingInfrastructureAkamaiAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAkamaiAPI {
	a := &EncodingInfrastructureAkamaiAPI{apiClient: apiClient}
	a.Regions = NewEncodingInfrastructureAkamaiRegionsAPIWithClient(apiClient)

	return a
}

// Create Add Akamai account
func (api *EncodingInfrastructureAkamaiAPI) Create(akamaiAccount model.AkamaiAccount) (*model.AkamaiAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AkamaiAccount
	err := api.apiClient.Post("/encoding/infrastructure/akamai", &akamaiAccount, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Akamai account
func (api *EncodingInfrastructureAkamaiAPI) Delete(infrastructureId string) (*model.AkamaiAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.AkamaiAccount
	err := api.apiClient.Delete("/encoding/infrastructure/akamai/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Akamai account details
func (api *EncodingInfrastructureAkamaiAPI) Get(infrastructureId string) (*model.AkamaiAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.AkamaiAccount
	err := api.apiClient.Get("/encoding/infrastructure/akamai/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Akamai accounts
func (api *EncodingInfrastructureAkamaiAPI) List(queryParams ...func(*EncodingInfrastructureAkamaiAPIListQueryParams)) (*pagination.AkamaiAccountsListPagination, error) {
	queryParameters := &EncodingInfrastructureAkamaiAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AkamaiAccountsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/akamai", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureAkamaiAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureAkamaiAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureAkamaiAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
