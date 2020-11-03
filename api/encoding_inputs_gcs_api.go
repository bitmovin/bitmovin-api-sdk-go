package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsGcsAPI communicates with '/encoding/inputs/gcs' endpoints
type EncodingInputsGcsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/gcs/{input_id}/customData' endpoints
    Customdata *EncodingInputsGcsCustomdataAPI

}

// NewEncodingInputsGcsAPI constructor for EncodingInputsGcsAPI that takes options as argument
func NewEncodingInputsGcsAPI(options ...apiclient.APIClientOption) (*EncodingInputsGcsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsGcsAPIWithClient(apiClient), nil
}

// NewEncodingInputsGcsAPIWithClient constructor for EncodingInputsGcsAPI that takes an APIClient as argument
func NewEncodingInputsGcsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsGcsAPI {
    a := &EncodingInputsGcsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsGcsCustomdataAPIWithClient(apiClient)

    return a
}

// Create GCS Input
func (api *EncodingInputsGcsAPI) Create(gcsInput model.GcsInput) (*model.GcsInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.GcsInput
    err := api.apiClient.Post("/encoding/inputs/gcs", &gcsInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete GCS Input
func (api *EncodingInputsGcsAPI) Delete(inputId string) (*model.GcsInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.GcsInput
    err := api.apiClient.Delete("/encoding/inputs/gcs/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get GCS Input Details
func (api *EncodingInputsGcsAPI) Get(inputId string) (*model.GcsInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.GcsInput
    err := api.apiClient.Get("/encoding/inputs/gcs/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List GCS Inputs
func (api *EncodingInputsGcsAPI) List(queryParams ...func(*EncodingInputsGcsAPIListQueryParams)) (*pagination.GcsInputsListPagination, error) {
    queryParameters := &EncodingInputsGcsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.GcsInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/gcs", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsGcsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsGcsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsGcsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


