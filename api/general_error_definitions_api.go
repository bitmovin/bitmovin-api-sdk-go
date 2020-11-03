package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// GeneralErrorDefinitionsAPI communicates with '/general/error-definitions' endpoints
type GeneralErrorDefinitionsAPI struct {
    apiClient *apiclient.APIClient
}

// NewGeneralErrorDefinitionsAPI constructor for GeneralErrorDefinitionsAPI that takes options as argument
func NewGeneralErrorDefinitionsAPI(options ...apiclient.APIClientOption) (*GeneralErrorDefinitionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewGeneralErrorDefinitionsAPIWithClient(apiClient), nil
}

// NewGeneralErrorDefinitionsAPIWithClient constructor for GeneralErrorDefinitionsAPI that takes an APIClient as argument
func NewGeneralErrorDefinitionsAPIWithClient(apiClient *apiclient.APIClient) *GeneralErrorDefinitionsAPI {
    a := &GeneralErrorDefinitionsAPI{apiClient: apiClient}
    return a
}

// List all possible api error definitions
func (api *GeneralErrorDefinitionsAPI) List(queryParams ...func(*GeneralErrorDefinitionsAPIListQueryParams)) (*pagination.ApiErrorDefinitionsListPagination, error) {
    queryParameters := &GeneralErrorDefinitionsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ApiErrorDefinitionsListPagination
    err := api.apiClient.Get("/general/error-definitions", nil, &responseModel, reqParams)
    return &responseModel, err
}

// GeneralErrorDefinitionsAPIListQueryParams contains all query parameters for the List endpoint
type GeneralErrorDefinitionsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *GeneralErrorDefinitionsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


