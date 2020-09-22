package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI communicates with '/notifications/webhooks/encoding/encodings/error/{webhook_id}/customData' endpoints
type NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsErrorCustomdataAPI constructor for NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsErrorCustomdataAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsErrorCustomdataAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsErrorCustomdataAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsErrorCustomdataAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI {
	a := &NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI{apiClient: apiClient}
	return a
}

// GetCustomDataByEncodingIdAndWebhookId Encoding Error Webhook Custom Data for specific Encoding Resource
func (api *NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/error/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetCustomDataByWebhookId Encoding Error Webhook Custom Data
func (api *NotificationsWebhooksEncodingEncodingsErrorCustomdataAPI) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/error/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
