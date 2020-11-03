package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsHttpsAPI communicates with '/encoding/inputs/https' endpoints
type EncodingInputsHttpsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/https/{input_id}/customData' endpoints
    Customdata *EncodingInputsHttpsCustomdataAPI

}

// NewEncodingInputsHttpsAPI constructor for EncodingInputsHttpsAPI that takes options as argument
func NewEncodingInputsHttpsAPI(options ...apiclient.APIClientOption) (*EncodingInputsHttpsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsHttpsAPIWithClient(apiClient), nil
}

// NewEncodingInputsHttpsAPIWithClient constructor for EncodingInputsHttpsAPI that takes an APIClient as argument
func NewEncodingInputsHttpsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsHttpsAPI {
    a := &EncodingInputsHttpsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsHttpsCustomdataAPIWithClient(apiClient)

    return a
}

// Create HTTPS Input
func (api *EncodingInputsHttpsAPI) Create(httpsInput model.HttpsInput) (*model.HttpsInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.HttpsInput
    err := api.apiClient.Post("/encoding/inputs/https", &httpsInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete HTTPS Input
func (api *EncodingInputsHttpsAPI) Delete(inputId string) (*model.HttpsInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.HttpsInput
    err := api.apiClient.Delete("/encoding/inputs/https/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get HTTPS Input Details
func (api *EncodingInputsHttpsAPI) Get(inputId string) (*model.HttpsInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.HttpsInput
    err := api.apiClient.Get("/encoding/inputs/https/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List HTTPS Inputs
func (api *EncodingInputsHttpsAPI) List(queryParams ...func(*EncodingInputsHttpsAPIListQueryParams)) (*pagination.HttpsInputsListPagination, error) {
    queryParameters := &EncodingInputsHttpsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.HttpsInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/https", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsHttpsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsHttpsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsHttpsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


