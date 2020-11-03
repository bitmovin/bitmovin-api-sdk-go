package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureAwsAPI communicates with '/encoding/infrastructure/aws' endpoints
type EncodingInfrastructureAwsAPI struct {
    apiClient *apiclient.APIClient

    // Regions communicates with '/encoding/infrastructure/aws/{infrastructure_id}/regions' endpoints
    Regions *EncodingInfrastructureAwsRegionsAPI

}

// NewEncodingInfrastructureAwsAPI constructor for EncodingInfrastructureAwsAPI that takes options as argument
func NewEncodingInfrastructureAwsAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAwsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInfrastructureAwsAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAwsAPIWithClient constructor for EncodingInfrastructureAwsAPI that takes an APIClient as argument
func NewEncodingInfrastructureAwsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAwsAPI {
    a := &EncodingInfrastructureAwsAPI{apiClient: apiClient}
    a.Regions = NewEncodingInfrastructureAwsRegionsAPIWithClient(apiClient)

    return a
}

// Create Add AWS Account
func (api *EncodingInfrastructureAwsAPI) Create(awsAccount model.AwsAccount) (*model.AwsAccount, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AwsAccount
    err := api.apiClient.Post("/encoding/infrastructure/aws", &awsAccount, &responseModel, reqParams)
    return &responseModel, err
}
// Delete AWS Account
func (api *EncodingInfrastructureAwsAPI) Delete(infrastructureId string) (*model.AwsAccount, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel model.AwsAccount
    err := api.apiClient.Delete("/encoding/infrastructure/aws/{infrastructure_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get AWS Account Details
func (api *EncodingInfrastructureAwsAPI) Get(infrastructureId string) (*model.AwsAccount, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel model.AwsAccount
    err := api.apiClient.Get("/encoding/infrastructure/aws/{infrastructure_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List AWS Accounts
func (api *EncodingInfrastructureAwsAPI) List(queryParams ...func(*EncodingInfrastructureAwsAPIListQueryParams)) (*pagination.AwsAccountsListPagination, error) {
    queryParameters := &EncodingInfrastructureAwsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AwsAccountsListPagination
    err := api.apiClient.Get("/encoding/infrastructure/aws", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInfrastructureAwsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureAwsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureAwsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


