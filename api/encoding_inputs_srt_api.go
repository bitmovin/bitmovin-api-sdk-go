package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsSrtAPI communicates with '/encoding/inputs/srt' endpoints
type EncodingInputsSrtAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/srt/{input_id}/customData' endpoints
    Customdata *EncodingInputsSrtCustomdataAPI

}

// NewEncodingInputsSrtAPI constructor for EncodingInputsSrtAPI that takes options as argument
func NewEncodingInputsSrtAPI(options ...apiclient.APIClientOption) (*EncodingInputsSrtAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsSrtAPIWithClient(apiClient), nil
}

// NewEncodingInputsSrtAPIWithClient constructor for EncodingInputsSrtAPI that takes an APIClient as argument
func NewEncodingInputsSrtAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsSrtAPI {
    a := &EncodingInputsSrtAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsSrtCustomdataAPIWithClient(apiClient)

    return a
}

// Create SRT input
func (api *EncodingInputsSrtAPI) Create(srtInput model.SrtInput) (*model.SrtInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.SrtInput
    err := api.apiClient.Post("/encoding/inputs/srt", &srtInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete SRT input
func (api *EncodingInputsSrtAPI) Delete(inputId string) (*model.SrtInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.SrtInput
    err := api.apiClient.Delete("/encoding/inputs/srt/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get SRT Input Details
func (api *EncodingInputsSrtAPI) Get(inputId string) (*model.SrtInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.SrtInput
    err := api.apiClient.Get("/encoding/inputs/srt/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List SRT inputs
func (api *EncodingInputsSrtAPI) List(queryParams ...func(*EncodingInputsSrtAPIListQueryParams)) (*pagination.SrtInputsListPagination, error) {
    queryParameters := &EncodingInputsSrtAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SrtInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/srt", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsSrtAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsSrtAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsSrtAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


