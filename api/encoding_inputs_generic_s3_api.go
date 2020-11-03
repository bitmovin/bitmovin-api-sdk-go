package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsGenericS3API communicates with '/encoding/inputs/generic-s3' endpoints
type EncodingInputsGenericS3API struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/generic-s3/{input_id}/customData' endpoints
    Customdata *EncodingInputsGenericS3CustomdataAPI

}

// NewEncodingInputsGenericS3API constructor for EncodingInputsGenericS3API that takes options as argument
func NewEncodingInputsGenericS3API(options ...apiclient.APIClientOption) (*EncodingInputsGenericS3API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsGenericS3APIWithClient(apiClient), nil
}

// NewEncodingInputsGenericS3APIWithClient constructor for EncodingInputsGenericS3API that takes an APIClient as argument
func NewEncodingInputsGenericS3APIWithClient(apiClient *apiclient.APIClient) *EncodingInputsGenericS3API {
    a := &EncodingInputsGenericS3API{apiClient: apiClient}
    a.Customdata = NewEncodingInputsGenericS3CustomdataAPIWithClient(apiClient)

    return a
}

// Create Generic S3 Input
func (api *EncodingInputsGenericS3API) Create(genericS3Input model.GenericS3Input) (*model.GenericS3Input, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.GenericS3Input
    err := api.apiClient.Post("/encoding/inputs/generic-s3", &genericS3Input, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Generic S3 Input
func (api *EncodingInputsGenericS3API) Delete(inputId string) (*model.GenericS3Input, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.GenericS3Input
    err := api.apiClient.Delete("/encoding/inputs/generic-s3/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Generic S3 Input Details
func (api *EncodingInputsGenericS3API) Get(inputId string) (*model.GenericS3Input, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.GenericS3Input
    err := api.apiClient.Get("/encoding/inputs/generic-s3/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Generic S3 Inputs
func (api *EncodingInputsGenericS3API) List(queryParams ...func(*EncodingInputsGenericS3APIListQueryParams)) (*pagination.GenericS3InputsListPagination, error) {
    queryParameters := &EncodingInputsGenericS3APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.GenericS3InputsListPagination
    err := api.apiClient.Get("/encoding/inputs/generic-s3", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsGenericS3APIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsGenericS3APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsGenericS3APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


