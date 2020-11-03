package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsGcsAPI communicates with '/encoding/outputs/gcs' endpoints
type EncodingOutputsGcsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/outputs/gcs/{output_id}/customData' endpoints
    Customdata *EncodingOutputsGcsCustomdataAPI

}

// NewEncodingOutputsGcsAPI constructor for EncodingOutputsGcsAPI that takes options as argument
func NewEncodingOutputsGcsAPI(options ...apiclient.APIClientOption) (*EncodingOutputsGcsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsGcsAPIWithClient(apiClient), nil
}

// NewEncodingOutputsGcsAPIWithClient constructor for EncodingOutputsGcsAPI that takes an APIClient as argument
func NewEncodingOutputsGcsAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsGcsAPI {
    a := &EncodingOutputsGcsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingOutputsGcsCustomdataAPIWithClient(apiClient)

    return a
}

// Create GCS Output
func (api *EncodingOutputsGcsAPI) Create(gcsOutput model.GcsOutput) (*model.GcsOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.GcsOutput
    err := api.apiClient.Post("/encoding/outputs/gcs", &gcsOutput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete GCS Output
func (api *EncodingOutputsGcsAPI) Delete(outputId string) (*model.GcsOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.GcsOutput
    err := api.apiClient.Delete("/encoding/outputs/gcs/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get GCS Output Details
func (api *EncodingOutputsGcsAPI) Get(outputId string) (*model.GcsOutput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.GcsOutput
    err := api.apiClient.Get("/encoding/outputs/gcs/{output_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List GCS Outputs
func (api *EncodingOutputsGcsAPI) List(queryParams ...func(*EncodingOutputsGcsAPIListQueryParams)) (*pagination.GcsOutputsListPagination, error) {
    queryParameters := &EncodingOutputsGcsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.GcsOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/gcs", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingOutputsGcsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsGcsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsGcsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


