package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsGcsServiceAccountAPI communicates with '/encoding/inputs/gcs-service-account' endpoints
type EncodingInputsGcsServiceAccountAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/gcs-service-account/{input_id}/customData' endpoints
    Customdata *EncodingInputsGcsServiceAccountCustomdataAPI

}

// NewEncodingInputsGcsServiceAccountAPI constructor for EncodingInputsGcsServiceAccountAPI that takes options as argument
func NewEncodingInputsGcsServiceAccountAPI(options ...apiclient.APIClientOption) (*EncodingInputsGcsServiceAccountAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsGcsServiceAccountAPIWithClient(apiClient), nil
}

// NewEncodingInputsGcsServiceAccountAPIWithClient constructor for EncodingInputsGcsServiceAccountAPI that takes an APIClient as argument
func NewEncodingInputsGcsServiceAccountAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsGcsServiceAccountAPI {
    a := &EncodingInputsGcsServiceAccountAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsGcsServiceAccountCustomdataAPIWithClient(apiClient)

    return a
}

// Create Service Account based GCS Input
func (api *EncodingInputsGcsServiceAccountAPI) Create(gcsServiceAccountInput model.GcsServiceAccountInput) (*model.GcsServiceAccountInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.GcsServiceAccountInput
    err := api.apiClient.Post("/encoding/inputs/gcs-service-account", &gcsServiceAccountInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Service Account based GCS Input
func (api *EncodingInputsGcsServiceAccountAPI) Delete(inputId string) (*model.GcsServiceAccountInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.GcsServiceAccountInput
    err := api.apiClient.Delete("/encoding/inputs/gcs-service-account/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get List Service Account based GCS Input Details
func (api *EncodingInputsGcsServiceAccountAPI) Get(inputId string) (*model.GcsServiceAccountInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.GcsServiceAccountInput
    err := api.apiClient.Get("/encoding/inputs/gcs-service-account/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Service Account based GCS Inputs
func (api *EncodingInputsGcsServiceAccountAPI) List(queryParams ...func(*EncodingInputsGcsServiceAccountAPIListQueryParams)) (*pagination.GcsServiceAccountInputsListPagination, error) {
    queryParameters := &EncodingInputsGcsServiceAccountAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.GcsServiceAccountInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/gcs-service-account", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsGcsServiceAccountAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsGcsServiceAccountAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsGcsServiceAccountAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


