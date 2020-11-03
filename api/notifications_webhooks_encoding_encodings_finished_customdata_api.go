package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI communicates with '/notifications/webhooks/encoding/encodings/finished/{webhook_id}/customData' endpoints
type NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI constructor for NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI {
    a := &NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI{apiClient: apiClient}
    return a
}

// GetCustomDataByEncodingIdAndWebhookId Encoding Finished Webhook Custom Data for specific Encoding Resource
func (api *NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/finished/{webhook_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}
// GetCustomDataByWebhookId Encoding Finished Webhook Custom Data
func (api *NotificationsWebhooksEncodingEncodingsFinishedCustomdataAPI) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/finished/{webhook_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

