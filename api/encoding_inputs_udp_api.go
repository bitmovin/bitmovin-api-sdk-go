package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsUdpAPI communicates with '/encoding/inputs/udp' endpoints
type EncodingInputsUdpAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsUdpAPI constructor for EncodingInputsUdpAPI that takes options as argument
func NewEncodingInputsUdpAPI(options ...apiclient.APIClientOption) (*EncodingInputsUdpAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsUdpAPIWithClient(apiClient), nil
}

// NewEncodingInputsUdpAPIWithClient constructor for EncodingInputsUdpAPI that takes an APIClient as argument
func NewEncodingInputsUdpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsUdpAPI {
    a := &EncodingInputsUdpAPI{apiClient: apiClient}
    return a
}

// Get UDP Input Details
func (api *EncodingInputsUdpAPI) Get(inputId string) (*model.UdpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.UdpInput
    err := api.apiClient.Get("/encoding/inputs/udp/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List UDP inputs
func (api *EncodingInputsUdpAPI) List(queryParams ...func(*EncodingInputsUdpAPIListQueryParams)) (*pagination.UdpInputsListPagination, error) {
    queryParameters := &EncodingInputsUdpAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.UdpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/udp", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsUdpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsUdpAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsUdpAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


