package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsTransferRetriesAPI communicates with '/encoding/encodings/{encoding_id}/transfer-retries' endpoints
type EncodingEncodingsTransferRetriesAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsTransferRetriesAPI constructor for EncodingEncodingsTransferRetriesAPI that takes options as argument
func NewEncodingEncodingsTransferRetriesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsTransferRetriesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsTransferRetriesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsTransferRetriesAPIWithClient constructor for EncodingEncodingsTransferRetriesAPI that takes an APIClient as argument
func NewEncodingEncodingsTransferRetriesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsTransferRetriesAPI {
    a := &EncodingEncodingsTransferRetriesAPI{apiClient: apiClient}
    return a
}

// Create Starts transfer retry. A transfer retry is only possible within 72 hours.
func (api *EncodingEncodingsTransferRetriesAPI) Create(encodingId string) (*model.TransferRetry, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.TransferRetry
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/transfer-retries", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Transfer retry Details
func (api *EncodingEncodingsTransferRetriesAPI) Get(encodingId string, transferRetryId string) (*model.TransferRetry, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["transfer_retry_id"] = transferRetryId
    }

    var responseModel model.TransferRetry
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/transfer-retries/{transfer_retry_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List transfer-retries
func (api *EncodingEncodingsTransferRetriesAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsTransferRetriesAPIListQueryParams)) (*pagination.TransferRetrysListPagination, error) {
    queryParameters := &EncodingEncodingsTransferRetriesAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.TransferRetrysListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/transfer-retries", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsTransferRetriesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsTransferRetriesAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsTransferRetriesAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


