package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// NotificationsWebhooksEncodingManifestErrorAPI communicates with '/notifications/webhooks/encoding/manifest/error' endpoints
type NotificationsWebhooksEncodingManifestErrorAPI struct {
    apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingManifestErrorAPI constructor for NotificationsWebhooksEncodingManifestErrorAPI that takes options as argument
func NewNotificationsWebhooksEncodingManifestErrorAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingManifestErrorAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsWebhooksEncodingManifestErrorAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingManifestErrorAPIWithClient constructor for NotificationsWebhooksEncodingManifestErrorAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingManifestErrorAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingManifestErrorAPI {
    a := &NotificationsWebhooksEncodingManifestErrorAPI{apiClient: apiClient}
    return a
}

// Create Add Manifest Error Webhook (All Manifests)
func (api *NotificationsWebhooksEncodingManifestErrorAPI) Create(webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/error", &webhook, &responseModel, reqParams)
    return &responseModel, err
}
// CreateByManifestId Add Manifest Error Webhook Notification (Specific Manifest)
func (api *NotificationsWebhooksEncodingManifestErrorAPI) CreateByManifestId(manifestId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/{manifest_id}/error", &webhook, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Manifest Error Webhook
func (api *NotificationsWebhooksEncodingManifestErrorAPI) Delete(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/manifest/error/{notification_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Update Replace Manifest Error Webhook Notification
func (api *NotificationsWebhooksEncodingManifestErrorAPI) Update(notificationId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel model.Webhook
    err := api.apiClient.Put("/notifications/webhooks/encoding/manifest/error/{notification_id}", &webhook, &responseModel, reqParams)
    return &responseModel, err
}

