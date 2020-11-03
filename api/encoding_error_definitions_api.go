package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingErrorDefinitionsAPI communicates with '/encoding/error-definitions' endpoints
type EncodingErrorDefinitionsAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingErrorDefinitionsAPI constructor for EncodingErrorDefinitionsAPI that takes options as argument
func NewEncodingErrorDefinitionsAPI(options ...apiclient.APIClientOption) (*EncodingErrorDefinitionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingErrorDefinitionsAPIWithClient(apiClient), nil
}

// NewEncodingErrorDefinitionsAPIWithClient constructor for EncodingErrorDefinitionsAPI that takes an APIClient as argument
func NewEncodingErrorDefinitionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingErrorDefinitionsAPI {
    a := &EncodingErrorDefinitionsAPI{apiClient: apiClient}
    return a
}

// List all possible encoding error definitions
func (api *EncodingErrorDefinitionsAPI) List(queryParams ...func(*EncodingErrorDefinitionsAPIListQueryParams)) (*pagination.EncodingErrorDefinitionsListPagination, error) {
    queryParameters := &EncodingErrorDefinitionsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.EncodingErrorDefinitionsListPagination
    err := api.apiClient.Get("/encoding/error-definitions", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingErrorDefinitionsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingErrorDefinitionsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingErrorDefinitionsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


