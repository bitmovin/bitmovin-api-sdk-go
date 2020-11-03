package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsS3RoleBasedAPI communicates with '/encoding/inputs/s3-role-based' endpoints
type EncodingInputsS3RoleBasedAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/s3-role-based/{input_id}/customData' endpoints
    Customdata *EncodingInputsS3RoleBasedCustomdataAPI

}

// NewEncodingInputsS3RoleBasedAPI constructor for EncodingInputsS3RoleBasedAPI that takes options as argument
func NewEncodingInputsS3RoleBasedAPI(options ...apiclient.APIClientOption) (*EncodingInputsS3RoleBasedAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsS3RoleBasedAPIWithClient(apiClient), nil
}

// NewEncodingInputsS3RoleBasedAPIWithClient constructor for EncodingInputsS3RoleBasedAPI that takes an APIClient as argument
func NewEncodingInputsS3RoleBasedAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsS3RoleBasedAPI {
    a := &EncodingInputsS3RoleBasedAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsS3RoleBasedCustomdataAPIWithClient(apiClient)

    return a
}

// Create S3 Role-based Input
func (api *EncodingInputsS3RoleBasedAPI) Create(s3RoleBasedInput model.S3RoleBasedInput) (*model.S3RoleBasedInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.S3RoleBasedInput
    err := api.apiClient.Post("/encoding/inputs/s3-role-based", &s3RoleBasedInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete S3 Role-based Input
func (api *EncodingInputsS3RoleBasedAPI) Delete(inputId string) (*model.S3RoleBasedInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.S3RoleBasedInput
    err := api.apiClient.Delete("/encoding/inputs/s3-role-based/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get S3 Role-based Input Details
func (api *EncodingInputsS3RoleBasedAPI) Get(inputId string) (*model.S3RoleBasedInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.S3RoleBasedInput
    err := api.apiClient.Get("/encoding/inputs/s3-role-based/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List S3 Role-based Inputs
func (api *EncodingInputsS3RoleBasedAPI) List(queryParams ...func(*EncodingInputsS3RoleBasedAPIListQueryParams)) (*pagination.S3RoleBasedInputsListPagination, error) {
    queryParameters := &EncodingInputsS3RoleBasedAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.S3RoleBasedInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/s3-role-based", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsS3RoleBasedAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsS3RoleBasedAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsS3RoleBasedAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


