package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsAzureAPI communicates with '/encoding/inputs/azure' endpoints
type EncodingInputsAzureAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/azure/{input_id}/customData' endpoints
    Customdata *EncodingInputsAzureCustomdataAPI

}

// NewEncodingInputsAzureAPI constructor for EncodingInputsAzureAPI that takes options as argument
func NewEncodingInputsAzureAPI(options ...apiclient.APIClientOption) (*EncodingInputsAzureAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsAzureAPIWithClient(apiClient), nil
}

// NewEncodingInputsAzureAPIWithClient constructor for EncodingInputsAzureAPI that takes an APIClient as argument
func NewEncodingInputsAzureAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAzureAPI {
    a := &EncodingInputsAzureAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsAzureCustomdataAPIWithClient(apiClient)

    return a
}

// Create Azure Input
func (api *EncodingInputsAzureAPI) Create(azureInput model.AzureInput) (*model.AzureInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AzureInput
    err := api.apiClient.Post("/encoding/inputs/azure", &azureInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Azure Input
func (api *EncodingInputsAzureAPI) Delete(inputId string) (*model.AzureInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.AzureInput
    err := api.apiClient.Delete("/encoding/inputs/azure/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Azure Input Details
func (api *EncodingInputsAzureAPI) Get(inputId string) (*model.AzureInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.AzureInput
    err := api.apiClient.Get("/encoding/inputs/azure/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Azure Inputs
func (api *EncodingInputsAzureAPI) List(queryParams ...func(*EncodingInputsAzureAPIListQueryParams)) (*pagination.AzureInputsListPagination, error) {
    queryParameters := &EncodingInputsAzureAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AzureInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/azure", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsAzureAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsAzureAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsAzureAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


