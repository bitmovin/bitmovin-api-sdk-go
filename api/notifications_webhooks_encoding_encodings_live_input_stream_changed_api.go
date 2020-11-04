package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI communicates with '/notifications/webhooks/encoding/encodings/live-input-stream-changed' endpoints
type NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI constructor for NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI {
	a := &NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI{apiClient: apiClient}
	return a
}

// Create Add Live Input Stream Changed Webhook Notification (All Encodings)
func (api *NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI) Create(webhookNotificationWithStreamConditionsRequest model.WebhookNotificationWithStreamConditionsRequest) (*model.WebhookNotificationWithStreamConditions, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.WebhookNotificationWithStreamConditions
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/live-input-stream-changed", &webhookNotificationWithStreamConditionsRequest, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByEncodingId Add Live Input Stream Changed Webhook Notification (Specific Encoding)
func (api *NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI) CreateByEncodingId(encodingId string, webhookNotificationWithStreamConditionsRequest model.WebhookNotificationWithStreamConditionsRequest) (*model.WebhookNotificationWithStreamConditions, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.WebhookNotificationWithStreamConditions
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/live-input-stream-changed", &webhookNotificationWithStreamConditionsRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Update Replace Live Input Stream Changed Webhook Notification
func (api *NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI) Update(notificationId string, webhookNotificationWithStreamConditionsRequest model.WebhookNotificationWithStreamConditionsRequest) (*model.WebhookNotificationWithStreamConditions, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["notification_id"] = notificationId
	}

	var responseModel model.WebhookNotificationWithStreamConditions
	err := api.apiClient.Put("/notifications/webhooks/encoding/encodings/live-input-stream-changed/{notification_id}", &webhookNotificationWithStreamConditionsRequest, &responseModel, reqParams)
	return &responseModel, err
}
