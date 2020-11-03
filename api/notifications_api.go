package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsAPI communicates with '/notifications' endpoints
type NotificationsAPI struct {
    apiClient *apiclient.APIClient

    // Webhooks intermediary API object with no endpoints
    Webhooks *NotificationsWebhooksAPI
    // States communicates with '/notifications/{notification_id}/states/{resource_id}' endpoints
    States *NotificationsStatesAPI
    // Emails communicates with '/notifications/emails' endpoints
    Emails *NotificationsEmailsAPI

}

// NewNotificationsAPI constructor for NotificationsAPI that takes options as argument
func NewNotificationsAPI(options ...apiclient.APIClientOption) (*NotificationsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsAPIWithClient(apiClient), nil
}

// NewNotificationsAPIWithClient constructor for NotificationsAPI that takes an APIClient as argument
func NewNotificationsAPIWithClient(apiClient *apiclient.APIClient) *NotificationsAPI {
    a := &NotificationsAPI{apiClient: apiClient}
    a.Webhooks = NewNotificationsWebhooksAPIWithClient(apiClient)
    a.States = NewNotificationsStatesAPIWithClient(apiClient)
    a.Emails = NewNotificationsEmailsAPIWithClient(apiClient)

    return a
}

// Delete Notification
func (api *NotificationsAPI) Delete(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/{notification_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Notification
func (api *NotificationsAPI) Get(notificationId string) (*model.Notification, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.Notification
    err := api.apiClient.Get("/notifications/{notification_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Notifications
func (api *NotificationsAPI) List(queryParams ...func(*NotificationsAPIListQueryParams)) (*pagination.NotificationsListPagination, error) {
    queryParameters := &NotificationsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.NotificationsListPagination
    err := api.apiClient.Get("/notifications", nil, &responseModel, reqParams)
    return &responseModel, err
}
// ListByNotificationId List Notification State History (All Resources)
func (api *NotificationsAPI) ListByNotificationId(notificationId string, queryParams ...func(*NotificationsAPIListByNotificationIdQueryParams)) (*pagination.NotificationStateEntrysListByNotificationIdPagination, error) {
    queryParameters := &NotificationsAPIListByNotificationIdQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.NotificationStateEntrysListByNotificationIdPagination
    err := api.apiClient.Get("/notifications/{notification_id}/states", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Mute Notification
func (api *NotificationsAPI) Mute(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Post("/notifications/{notification_id}/mute", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Unmute Notification
func (api *NotificationsAPI) Unmute(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Post("/notifications/{notification_id}/unmute", nil, &responseModel, reqParams)
    return &responseModel, err
}

// NotificationsAPIListQueryParams contains all query parameters for the List endpoint
type NotificationsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


// NotificationsAPIListByNotificationIdQueryParams contains all query parameters for the ListByNotificationId endpoint
type NotificationsAPIListByNotificationIdQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *NotificationsAPIListByNotificationIdQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


