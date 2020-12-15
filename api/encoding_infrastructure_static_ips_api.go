package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureStaticIpsAPI communicates with '/encoding/infrastructure/static-ips' endpoints
type EncodingInfrastructureStaticIpsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureStaticIpsAPI constructor for EncodingInfrastructureStaticIpsAPI that takes options as argument
func NewEncodingInfrastructureStaticIpsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureStaticIpsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureStaticIpsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureStaticIpsAPIWithClient constructor for EncodingInfrastructureStaticIpsAPI that takes an APIClient as argument
func NewEncodingInfrastructureStaticIpsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureStaticIpsAPI {
	a := &EncodingInfrastructureStaticIpsAPI{apiClient: apiClient}
	return a
}

// Create Add Static IP Address
func (api *EncodingInfrastructureStaticIpsAPI) Create(staticIp model.StaticIp) (*model.StaticIp, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StaticIp
	err := api.apiClient.Post("/encoding/infrastructure/static-ips", &staticIp, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Static IP Address
func (api *EncodingInfrastructureStaticIpsAPI) Delete(id string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["id"] = id
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/infrastructure/static-ips/{id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Static IP Address Details
func (api *EncodingInfrastructureStaticIpsAPI) Get(id string) (*model.StaticIp, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["id"] = id
	}

	var responseModel model.StaticIp
	err := api.apiClient.Get("/encoding/infrastructure/static-ips/{id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Static IP Addresses
func (api *EncodingInfrastructureStaticIpsAPI) List(queryParams ...func(*EncodingInfrastructureStaticIpsAPIListQueryParams)) (*pagination.StaticIpsListPagination, error) {
	queryParameters := &EncodingInfrastructureStaticIpsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StaticIpsListPagination
	err := api.apiClient.Get("/encoding/infrastructure/static-ips", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureStaticIpsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureStaticIpsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureStaticIpsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
