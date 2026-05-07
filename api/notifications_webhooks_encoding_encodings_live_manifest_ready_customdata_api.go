package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI communicates with '/notifications/webhooks/encoding/encodings/live-manifest-ready/{webhook_id}/customData' endpoints
type NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI constructor for NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI {
	a := &NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI{apiClient: apiClient}
	return a
}

// GetCustomDataByEncodingIdAndWebhookId &#39;Live Manifest Ready&#39; Webhook Custom Data for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/live-manifest-ready/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetCustomDataByWebhookId &#39;Live Manifest Ready&#39; Webhook Custom Data
func (api *NotificationsWebhooksEncodingEncodingsLiveManifestReadyCustomdataAPI) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/live-manifest-ready/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
