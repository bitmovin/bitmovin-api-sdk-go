package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsEmailsAPI communicates with '/notifications/emails' endpoints
type NotificationsEmailsAPI struct {
    apiClient *apiclient.APIClient

    // UsageReports communicates with '/notifications/emails/usage-reports' endpoints
    UsageReports *NotificationsEmailsUsageReportsAPI
    // Encoding communicates with '/notifications/emails/encoding' endpoints
    Encoding *NotificationsEmailsEncodingAPI

}

// NewNotificationsEmailsAPI constructor for NotificationsEmailsAPI that takes options as argument
func NewNotificationsEmailsAPI(options ...apiclient.APIClientOption) (*NotificationsEmailsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsEmailsAPIWithClient(apiClient), nil
}

// NewNotificationsEmailsAPIWithClient constructor for NotificationsEmailsAPI that takes an APIClient as argument
func NewNotificationsEmailsAPIWithClient(apiClient *apiclient.APIClient) *NotificationsEmailsAPI {
    a := &NotificationsEmailsAPI{apiClient: apiClient}
    a.UsageReports = NewNotificationsEmailsUsageReportsAPIWithClient(apiClient)
    a.Encoding = NewNotificationsEmailsEncodingAPIWithClient(apiClient)

    return a
}

// List Email Notifications
func (api *NotificationsEmailsAPI) List(queryParams ...func(*NotificationsEmailsAPIListQueryParams)) (*pagination.NotificationsListPagination, error) {
    queryParameters := &NotificationsEmailsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.NotificationsListPagination
    err := api.apiClient.Get("/notifications/emails", nil, &responseModel, reqParams)
    return &responseModel, err
}

// NotificationsEmailsAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsEmailsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsEmailsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


