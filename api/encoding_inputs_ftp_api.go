package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsFtpAPI communicates with '/encoding/inputs/ftp' endpoints
type EncodingInputsFtpAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/inputs/ftp/{input_id}/customData' endpoints
    Customdata *EncodingInputsFtpCustomdataAPI

}

// NewEncodingInputsFtpAPI constructor for EncodingInputsFtpAPI that takes options as argument
func NewEncodingInputsFtpAPI(options ...apiclient.APIClientOption) (*EncodingInputsFtpAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsFtpAPIWithClient(apiClient), nil
}

// NewEncodingInputsFtpAPIWithClient constructor for EncodingInputsFtpAPI that takes an APIClient as argument
func NewEncodingInputsFtpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsFtpAPI {
    a := &EncodingInputsFtpAPI{apiClient: apiClient}
    a.Customdata = NewEncodingInputsFtpCustomdataAPIWithClient(apiClient)

    return a
}

// Create FTP Input
func (api *EncodingInputsFtpAPI) Create(ftpInput model.FtpInput) (*model.FtpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.FtpInput
    err := api.apiClient.Post("/encoding/inputs/ftp", &ftpInput, &responseModel, reqParams)
    return &responseModel, err
}
// Delete FTP Input
func (api *EncodingInputsFtpAPI) Delete(inputId string) (*model.FtpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.FtpInput
    err := api.apiClient.Delete("/encoding/inputs/ftp/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get FTP Input Details
func (api *EncodingInputsFtpAPI) Get(inputId string) (*model.FtpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.FtpInput
    err := api.apiClient.Get("/encoding/inputs/ftp/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List FTP Inputs
func (api *EncodingInputsFtpAPI) List(queryParams ...func(*EncodingInputsFtpAPIListQueryParams)) (*pagination.FtpInputsListPagination, error) {
    queryParameters := &EncodingInputsFtpAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.FtpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/ftp", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsFtpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsFtpAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsFtpAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


