package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI communicates with '/notifications/webhooks/encoding/encodings/encoding-status-changed' endpoints
type NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI constructor for NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI {
	a := &NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI{apiClient: apiClient}
	return a
}

// Create Add Encoding Changed Webhook Notification (All Encodings)
func (api *NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI) Create(webhookNotificationWithStreamConditionsRequest model.WebhookNotificationWithStreamConditionsRequest) (*model.WebhookNotificationWithStreamConditions, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.WebhookNotificationWithStreamConditions
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/encoding-status-changed", &webhookNotificationWithStreamConditionsRequest, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByEncodingId Add Encoding Changed Webhook Notification (Specific Encoding)
func (api *NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI) CreateByEncodingId(encodingId string, webhookNotificationWithStreamConditionsRequest model.WebhookNotificationWithStreamConditionsRequest) (*model.WebhookNotificationWithStreamConditions, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.WebhookNotificationWithStreamConditions
	err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/encoding-status-changed", &webhookNotificationWithStreamConditionsRequest, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteByWebhookId Delete Encoding Status Changed Webhook
func (api *NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI) DeleteByWebhookId(notificationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["notification_id"] = notificationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/encoding-status-changed/{notification_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update Replace Encoding Status Changed Webhook Notification
func (api *NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI) Update(notificationId string, webhookNotificationWithStreamConditionsRequest model.WebhookNotificationWithStreamConditionsRequest) (*model.WebhookNotificationWithStreamConditions, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["notification_id"] = notificationId
	}

	var responseModel model.WebhookNotificationWithStreamConditions
	err := api.apiClient.Put("/notifications/webhooks/encoding/encodings/encoding-status-changed/{notification_id}", &webhookNotificationWithStreamConditionsRequest, &responseModel, reqParams)
	return &responseModel, err
}
