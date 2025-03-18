package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureOciAPI communicates with '/encoding/infrastructure/oci' endpoints
type EncodingInfrastructureOciAPI struct {
	apiClient *apiclient.APIClient

	// Regions communicates with '/encoding/infrastructure/oci/{infrastructure_id}/regions' endpoints
	Regions *EncodingInfrastructureOciRegionsAPI
}

// NewEncodingInfrastructureOciAPI constructor for EncodingInfrastructureOciAPI that takes options as argument
func NewEncodingInfrastructureOciAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureOciAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureOciAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureOciAPIWithClient constructor for EncodingInfrastructureOciAPI that takes an APIClient as argument
func NewEncodingInfrastructureOciAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureOciAPI {
	a := &EncodingInfrastructureOciAPI{apiClient: apiClient}
	a.Regions = NewEncodingInfrastructureOciRegionsAPIWithClient(apiClient)

	return a
}

// Create Add OCI account
func (api *EncodingInfrastructureOciAPI) Create(ociAccount model.OciAccount) (*model.OciAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.OciAccount
	err := api.apiClient.Post("/encoding/infrastructure/oci", &ociAccount, &responseModel, reqParams)
	return &responseModel, err
}

// Delete OCI account
func (api *EncodingInfrastructureOciAPI) Delete(infrastructureId string) (*model.OciAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.OciAccount
	err := api.apiClient.Delete("/encoding/infrastructure/oci/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get OCI account details
func (api *EncodingInfrastructureOciAPI) Get(infrastructureId string) (*model.OciAccount, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.OciAccount
	err := api.apiClient.Get("/encoding/infrastructure/oci/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List OCI accounts
func (api *EncodingInfrastructureOciAPI) List(queryParams ...func(*EncodingInfrastructureOciAPIListQueryParams)) (*pagination.OciAccountsListPagination, error) {
	queryParameters := &EncodingInfrastructureOciAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.OciAccountsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/oci", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureOciAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureOciAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureOciAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
