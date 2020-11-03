package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsLiveInsertableContentAPI communicates with '/encoding/encodings/{encoding_id}/live/insertable-content' endpoints
type EncodingEncodingsLiveInsertableContentAPI struct {
    apiClient *apiclient.APIClient

    // Schedule communicates with '/encoding/encodings/{encoding_id}/live/insertable-content/{content_id}/schedule' endpoints
    Schedule *EncodingEncodingsLiveInsertableContentScheduleAPI
    // Scheduled communicates with '/encoding/encodings/{encoding_id}/live/insertable-content/scheduled' endpoints
    Scheduled *EncodingEncodingsLiveInsertableContentScheduledAPI
    // Stop communicates with '/encoding/encodings/{encoding_id}/live/insertable-content/stop' endpoints
    Stop *EncodingEncodingsLiveInsertableContentStopAPI

}

// NewEncodingEncodingsLiveInsertableContentAPI constructor for EncodingEncodingsLiveInsertableContentAPI that takes options as argument
func NewEncodingEncodingsLiveInsertableContentAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveInsertableContentAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsLiveInsertableContentAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveInsertableContentAPIWithClient constructor for EncodingEncodingsLiveInsertableContentAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveInsertableContentAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveInsertableContentAPI {
    a := &EncodingEncodingsLiveInsertableContentAPI{apiClient: apiClient}
    a.Schedule = NewEncodingEncodingsLiveInsertableContentScheduleAPIWithClient(apiClient)
    a.Scheduled = NewEncodingEncodingsLiveInsertableContentScheduledAPIWithClient(apiClient)
    a.Stop = NewEncodingEncodingsLiveInsertableContentStopAPIWithClient(apiClient)

    return a
}

// Create Make Insertable Content Available For A Live Encoding
func (api *EncodingEncodingsLiveInsertableContentAPI) Create(encodingId string, insertableContent model.InsertableContent) (*model.InsertableContent, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.InsertableContent
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/insertable-content", &insertableContent, &responseModel, reqParams)
    return &responseModel, err
}
// List All Insertable Content Available For A Live Encoding
func (api *EncodingEncodingsLiveInsertableContentAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsLiveInsertableContentAPIListQueryParams)) (*pagination.InsertableContentsListPagination, error) {
    queryParameters := &EncodingEncodingsLiveInsertableContentAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.InsertableContentsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/insertable-content", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsLiveInsertableContentAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsLiveInsertableContentAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsLiveInsertableContentAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


