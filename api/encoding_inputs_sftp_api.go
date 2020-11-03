package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsSftpAPI communicates with '/encoding/inputs/sftp' endpoints
type EncodingInputsSftpAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/sftp/{input_id}/customData' endpoints
    Customdata *EncodingInputsSftpCustomdataAPI

}

// NewEncodingInputsSftpAPI constructor for EncodingInputsSftpAPI that takes options as argument
func NewEncodingInputsSftpAPI(options ...apiclient.APIClientOption) (*EncodingInputsSftpAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsSftpAPIWithClient(apiClient), nil
}

// NewEncodingInputsSftpAPIWithClient constructor for EncodingInputsSftpAPI that takes an APIClient as argument
func NewEncodingInputsSftpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsSftpAPI {
    a := &EncodingInputsSftpAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsSftpCustomdataAPIWithClient(apiClient)

    return a
}

// Create SFTP Input
func (api *EncodingInputsSftpAPI) Create(sftpInput model.SftpInput) (*model.SftpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.SftpInput
    err := api.apiClient.Post("/encoding/inputs/sftp", &sftpInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete SFTP Input
func (api *EncodingInputsSftpAPI) Delete(inputId string) (*model.SftpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.SftpInput
    err := api.apiClient.Delete("/encoding/inputs/sftp/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get SFTP Input Details
func (api *EncodingInputsSftpAPI) Get(inputId string) (*model.SftpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.SftpInput
    err := api.apiClient.Get("/encoding/inputs/sftp/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List SFTP Inputs
func (api *EncodingInputsSftpAPI) List(queryParams ...func(*EncodingInputsSftpAPIListQueryParams)) (*pagination.SftpInputsListPagination, error) {
    queryParameters := &EncodingInputsSftpAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SftpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/sftp", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsSftpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsSftpAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsSftpAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


