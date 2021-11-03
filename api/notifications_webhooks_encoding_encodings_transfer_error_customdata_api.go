package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI communicates with '/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}/customData' endpoints
type NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI constructor for NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI {
	a := &NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI{apiClient: apiClient}
	return a
}

// GetCustomDataByEncodingIdAndWebhookId &#39;Encoding Transfer Error&#39; Webhook Custom Data for a specific Encoding
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/transfer-error/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetCustomDataByWebhookId &#39;Encoding Transfer Error&#39; Webhook Custom Data
func (api *NotificationsWebhooksEncodingEncodingsTransferErrorCustomdataAPI) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["webhook_id"] = webhookId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/transfer-error/{webhook_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
