package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI communicates with '/notifications/webhooks/encoding/encodings/live-encoding-heartbeat/{webhook_id}/customData' endpoints
type NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI constructor for NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI {
	a := &NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI{apiClient: apiClient}
	return a
}

// GetCustomDataByEncodingIdAndWebhookId Get &#39;Live Encoding Heartbeat&#39; Webhook Custom Data for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/live-encoding-heartbeat/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetCustomDataByWebhookId Get &#39;Live Encoding Heartbeat&#39; Webhook Custom Data
func (api *NotificationsWebhooksEncodingEncodingsLiveEncodingHeartbeatCustomdataAPI) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/live-encoding-heartbeat/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
