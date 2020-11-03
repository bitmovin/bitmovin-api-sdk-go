package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsAsperaAPI communicates with '/encoding/inputs/aspera' endpoints
type EncodingInputsAsperaAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/aspera/{input_id}/customData' endpoints
    Customdata *EncodingInputsAsperaCustomdataAPI

}

// NewEncodingInputsAsperaAPI constructor for EncodingInputsAsperaAPI that takes options as argument
func NewEncodingInputsAsperaAPI(options ...apiclient.APIClientOption) (*EncodingInputsAsperaAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsAsperaAPIWithClient(apiClient), nil
}

// NewEncodingInputsAsperaAPIWithClient constructor for EncodingInputsAsperaAPI that takes an APIClient as argument
func NewEncodingInputsAsperaAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAsperaAPI {
    a := &EncodingInputsAsperaAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsAsperaCustomdataAPIWithClient(apiClient)

    return a
}

// Create Aspera Input
func (api *EncodingInputsAsperaAPI) Create(asperaInput model.AsperaInput) (*model.AsperaInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AsperaInput
    err := api.apiClient.Post("/encoding/inputs/aspera", &asperaInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Aspera Input
func (api *EncodingInputsAsperaAPI) Delete(inputId string) (*model.AsperaInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.AsperaInput
    err := api.apiClient.Delete("/encoding/inputs/aspera/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Aspera Input Details
func (api *EncodingInputsAsperaAPI) Get(inputId string) (*model.AsperaInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.AsperaInput
    err := api.apiClient.Get("/encoding/inputs/aspera/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Aspera Inputs
func (api *EncodingInputsAsperaAPI) List(queryParams ...func(*EncodingInputsAsperaAPIListQueryParams)) (*pagination.AsperaInputsListPagination, error) {
    queryParameters := &EncodingInputsAsperaAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AsperaInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/aspera", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsAsperaAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsAsperaAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsAsperaAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


