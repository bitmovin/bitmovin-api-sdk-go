package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI communicates with '/notifications/emails/encoding/encodings/live-input-stream-changed' endpoints
type NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI struct {
    apiClient *apiclient.APIClient
}

// NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI constructor for NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI that takes options as argument
func NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI(options ...apiclient.APIClientOption) (*NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPIWithClient(apiClient), nil
}

// NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPIWithClient constructor for NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI that takes an APIClient as argument
func NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPIWithClient(apiClient *apiclient.APIClient) *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI {
    a := &NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI{apiClient: apiClient}
    return a
}

// Create Add Live Input Stream Changed Email Notification (All Encodings)
func (api *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI) Create(emailNotificationWithStreamConditionsRequest model.EmailNotificationWithStreamConditionsRequest) (*model.EmailNotificationWithStreamConditions, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.EmailNotificationWithStreamConditions
    err := api.apiClient.Post("/notifications/emails/encoding/encodings/live-input-stream-changed", &emailNotificationWithStreamConditionsRequest, &responseModel, reqParams)
    return &responseModel, err
}
// CreateByEncodingId Add Live Input Stream Changed Email Notification (Specific Encoding)
func (api *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI) CreateByEncodingId(encodingId string, emailNotificationWithStreamConditionsRequest model.EmailNotificationWithStreamConditionsRequest) (*model.EmailNotificationWithStreamConditions, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.EmailNotificationWithStreamConditions
    err := api.apiClient.Post("/notifications/emails/encoding/encodings/{encoding_id}/live-input-stream-changed", &emailNotificationWithStreamConditionsRequest, &responseModel, reqParams)
    return &responseModel, err
}
// Update Replace Live Input Stream Changed Email Notification
func (api *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedAPI) Update(notificationId string, emailNotificationWithStreamConditionsRequest model.EmailNotificationWithStreamConditionsRequest) (*model.EmailNotificationWithStreamConditions, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.EmailNotificationWithStreamConditions
    err := api.apiClient.Put("/notifications/emails/encoding/encodings/live-input-stream-changed/{notification_id}", &emailNotificationWithStreamConditionsRequest, &responseModel, reqParams)
    return &responseModel, err
}

