package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsEmailsEncodingAPI communicates with '/notifications/emails/encoding' endpoints
type NotificationsEmailsEncodingAPI struct {
    apiClient *apiclient.APIClient

    // Encodings communicates with '/notifications/emails/encoding/encodings/{encoding_id}' endpoints
    Encodings *NotificationsEmailsEncodingEncodingsAPI

}

// NewNotificationsEmailsEncodingAPI constructor for NotificationsEmailsEncodingAPI that takes options as argument
func NewNotificationsEmailsEncodingAPI(options ...apiclient.APIClientOption) (*NotificationsEmailsEncodingAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsEmailsEncodingAPIWithClient(apiClient), nil
}

// NewNotificationsEmailsEncodingAPIWithClient constructor for NotificationsEmailsEncodingAPI that takes an APIClient as argument
func NewNotificationsEmailsEncodingAPIWithClient(apiClient *apiclient.APIClient) *NotificationsEmailsEncodingAPI {
    a := &NotificationsEmailsEncodingAPI{apiClient: apiClient}
    a.Encodings = NewNotificationsEmailsEncodingEncodingsAPIWithClient(apiClient)

    return a
}

// List Email Notifications (All Encodings)
func (api *NotificationsEmailsEncodingAPI) List(queryParams ...func(*NotificationsEmailsEncodingAPIListQueryParams)) (*pagination.EmailNotificationsListPagination, error) {
    queryParameters := &NotificationsEmailsEncodingAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.EmailNotificationsListPagination
    err := api.apiClient.Get("/notifications/emails/encoding", nil, &responseModel, reqParams)
    return &responseModel, err
}

// NotificationsEmailsEncodingAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsEmailsEncodingAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsEmailsEncodingAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


