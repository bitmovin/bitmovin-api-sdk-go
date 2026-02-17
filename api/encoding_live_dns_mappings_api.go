package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingLiveDnsMappingsAPI communicates with '/encoding/live/dns-mappings' endpoints
type EncodingLiveDnsMappingsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveDnsMappingsAPI constructor for EncodingLiveDnsMappingsAPI that takes options as argument
func NewEncodingLiveDnsMappingsAPI(options ...apiclient.APIClientOption) (*EncodingLiveDnsMappingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveDnsMappingsAPIWithClient(apiClient), nil
}

// NewEncodingLiveDnsMappingsAPIWithClient constructor for EncodingLiveDnsMappingsAPI that takes an APIClient as argument
func NewEncodingLiveDnsMappingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveDnsMappingsAPI {
	a := &EncodingLiveDnsMappingsAPI{apiClient: apiClient}
	return a
}

// List DNS Mappings
func (api *EncodingLiveDnsMappingsAPI) List(queryParams ...func(*EncodingLiveDnsMappingsAPIListQueryParams)) (*pagination.DnsMappingResponsesListPagination, error) {
	queryParameters := &EncodingLiveDnsMappingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DnsMappingResponsesListPagination
	err := api.apiClient.Get("/encoding/live/dns-mappings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingLiveDnsMappingsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingLiveDnsMappingsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingLiveDnsMappingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
