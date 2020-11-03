package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsLiveInsertableContentScheduledAPI communicates with '/encoding/encodings/{encoding_id}/live/insertable-content/scheduled' endpoints
type EncodingEncodingsLiveInsertableContentScheduledAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveInsertableContentScheduledAPI constructor for EncodingEncodingsLiveInsertableContentScheduledAPI that takes options as argument
func NewEncodingEncodingsLiveInsertableContentScheduledAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveInsertableContentScheduledAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsLiveInsertableContentScheduledAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveInsertableContentScheduledAPIWithClient constructor for EncodingEncodingsLiveInsertableContentScheduledAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveInsertableContentScheduledAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveInsertableContentScheduledAPI {
    a := &EncodingEncodingsLiveInsertableContentScheduledAPI{apiClient: apiClient}
    return a
}

// List All Scheduled Insertable Content For A Live Encoding
func (api *EncodingEncodingsLiveInsertableContentScheduledAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsLiveInsertableContentScheduledAPIListQueryParams)) (*pagination.ScheduledInsertableContentsListPagination, error) {
    queryParameters := &EncodingEncodingsLiveInsertableContentScheduledAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ScheduledInsertableContentsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/insertable-content/scheduled", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsLiveInsertableContentScheduledAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsLiveInsertableContentScheduledAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsLiveInsertableContentScheduledAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


