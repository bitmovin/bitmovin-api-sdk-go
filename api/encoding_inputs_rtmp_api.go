package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsRtmpAPI communicates with '/encoding/inputs/rtmp' endpoints
type EncodingInputsRtmpAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsRtmpAPI constructor for EncodingInputsRtmpAPI that takes options as argument
func NewEncodingInputsRtmpAPI(options ...apiclient.APIClientOption) (*EncodingInputsRtmpAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsRtmpAPIWithClient(apiClient), nil
}

// NewEncodingInputsRtmpAPIWithClient constructor for EncodingInputsRtmpAPI that takes an APIClient as argument
func NewEncodingInputsRtmpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsRtmpAPI {
    a := &EncodingInputsRtmpAPI{apiClient: apiClient}
    return a
}

// Get RTMP Input Details
func (api *EncodingInputsRtmpAPI) Get(inputId string) (*model.RtmpInput, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.RtmpInput
    err := api.apiClient.Get("/encoding/inputs/rtmp/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List RTMP Inputs
func (api *EncodingInputsRtmpAPI) List(queryParams ...func(*EncodingInputsRtmpAPIListQueryParams)) (*pagination.RtmpInputsListPagination, error) {
    queryParameters := &EncodingInputsRtmpAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.RtmpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/rtmp", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsRtmpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsRtmpAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsRtmpAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


