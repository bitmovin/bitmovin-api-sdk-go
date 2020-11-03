package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsAzureAPI communicates with '/encoding/outputs/azure' endpoints
type EncodingOutputsAzureAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/outputs/azure/{output_id}/customData' endpoints
    Customdata *EncodingOutputsAzureCustomdataAPI

}

// NewEncodingOutputsAzureAPI constructor for EncodingOutputsAzureAPI that takes options as argument
func NewEncodingOutputsAzureAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAzureAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsAzureAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAzureAPIWithClient constructor for EncodingOutputsAzureAPI that takes an APIClient as argument
func NewEncodingOutputsAzureAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAzureAPI {
    a := &EncodingOutputsAzureAPI{apiClient: apiClient}
    a.Customdata = NewEncodingOutputsAzureCustomdataAPIWithClient(apiClient)

    return a
}

// Create Azure Output
func (api *EncodingOutputsAzureAPI) Create(azureOutput model.AzureOutput) (*model.AzureOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AzureOutput
    err := api.apiClient.Post("/encoding/outputs/azure", &azureOutput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Azure Output
func (api *EncodingOutputsAzureAPI) Delete(outputId string) (*model.AzureOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.AzureOutput
    err := api.apiClient.Delete("/encoding/outputs/azure/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Azure Output Details
func (api *EncodingOutputsAzureAPI) Get(outputId string) (*model.AzureOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.AzureOutput
    err := api.apiClient.Get("/encoding/outputs/azure/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Azure Outputs
func (api *EncodingOutputsAzureAPI) List(queryParams ...func(*EncodingOutputsAzureAPIListQueryParams)) (*pagination.AzureOutputsListPagination, error) {
    queryParameters := &EncodingOutputsAzureAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AzureOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/azure", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingOutputsAzureAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsAzureAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsAzureAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


