package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsEmailsEncodingEncodingsErrorAPI communicates with '/notifications/emails/encoding/encodings/error' endpoints
type NotificationsEmailsEncodingEncodingsErrorAPI struct {
    apiClient *apiclient.APIClient
}

// NewNotificationsEmailsEncodingEncodingsErrorAPI constructor for NotificationsEmailsEncodingEncodingsErrorAPI that takes options as argument
func NewNotificationsEmailsEncodingEncodingsErrorAPI(options ...apiclient.APIClientOption) (*NotificationsEmailsEncodingEncodingsErrorAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsEmailsEncodingEncodingsErrorAPIWithClient(apiClient), nil
}

// NewNotificationsEmailsEncodingEncodingsErrorAPIWithClient constructor for NotificationsEmailsEncodingEncodingsErrorAPI that takes an APIClient as argument
func NewNotificationsEmailsEncodingEncodingsErrorAPIWithClient(apiClient *apiclient.APIClient) *NotificationsEmailsEncodingEncodingsErrorAPI {
    a := &NotificationsEmailsEncodingEncodingsErrorAPI{apiClient: apiClient}
    return a
}

// CreateByEncodingId Add Encoding Error Email Notification (Specific Encoding)
func (api *NotificationsEmailsEncodingEncodingsErrorAPI) CreateByEncodingId(encodingId string, emailNotification model.EmailNotification) (*model.EmailNotification, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.EmailNotification
    err := api.apiClient.Post("/notifications/emails/encoding/encodings/{encoding_id}/error", &emailNotification, &responseModel, reqParams)
    return &responseModel, err
}
// Update Replace Encoding Error Email Notification
func (api *NotificationsEmailsEncodingEncodingsErrorAPI) Update(notificationId string, emailNotification model.EmailNotification) (*model.EmailNotification, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.EmailNotification
    err := api.apiClient.Put("/notifications/emails/encoding/encodings/error/{notification_id}", &emailNotification, &responseModel, reqParams)
    return &responseModel, err
}

