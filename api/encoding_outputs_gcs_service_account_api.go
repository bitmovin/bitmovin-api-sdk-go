package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsGcsServiceAccountAPI communicates with '/encoding/outputs/gcs-service-account' endpoints
type EncodingOutputsGcsServiceAccountAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/outputs/gcs-service-account/{output_id}/customData' endpoints
    Customdata *EncodingOutputsGcsServiceAccountCustomdataAPI

}

// NewEncodingOutputsGcsServiceAccountAPI constructor for EncodingOutputsGcsServiceAccountAPI that takes options as argument
func NewEncodingOutputsGcsServiceAccountAPI(options ...apiclient.APIClientOption) (*EncodingOutputsGcsServiceAccountAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsGcsServiceAccountAPIWithClient(apiClient), nil
}

// NewEncodingOutputsGcsServiceAccountAPIWithClient constructor for EncodingOutputsGcsServiceAccountAPI that takes an APIClient as argument
func NewEncodingOutputsGcsServiceAccountAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsGcsServiceAccountAPI {
    a := &EncodingOutputsGcsServiceAccountAPI{apiClient: apiClient}
    a.Customdata = NewEncodingOutputsGcsServiceAccountCustomdataAPIWithClient(apiClient)

    return a
}

// Create Service Account based GCS Output
func (api *EncodingOutputsGcsServiceAccountAPI) Create(gcsServiceAccountOutput model.GcsServiceAccountOutput) (*model.GcsServiceAccountOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.GcsServiceAccountOutput
    err := api.apiClient.Post("/encoding/outputs/gcs-service-account", &gcsServiceAccountOutput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Service Account based GCS Output
func (api *EncodingOutputsGcsServiceAccountAPI) Delete(outputId string) (*model.GcsServiceAccountOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.GcsServiceAccountOutput
    err := api.apiClient.Delete("/encoding/outputs/gcs-service-account/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Service Account based GCS Output Details
func (api *EncodingOutputsGcsServiceAccountAPI) Get(outputId string) (*model.GcsServiceAccountOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.GcsServiceAccountOutput
    err := api.apiClient.Get("/encoding/outputs/gcs-service-account/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Service Account based GCS Outputs
func (api *EncodingOutputsGcsServiceAccountAPI) List(queryParams ...func(*EncodingOutputsGcsServiceAccountAPIListQueryParams)) (*pagination.GcsServiceAccountOutputsListPagination, error) {
    queryParameters := &EncodingOutputsGcsServiceAccountAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.GcsServiceAccountOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/gcs-service-account", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingOutputsGcsServiceAccountAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsGcsServiceAccountAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsGcsServiceAccountAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


