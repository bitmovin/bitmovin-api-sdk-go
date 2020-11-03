package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// NotificationsWebhooksEncodingAPI intermediary API object with no endpoints
type NotificationsWebhooksEncodingAPI struct {
    apiClient *apiclient.APIClient

    // Encodings intermediary API object with no endpoints
    Encodings *NotificationsWebhooksEncodingEncodingsAPI
    // Manifest communicates with '/notifications/webhooks/encoding/manifest/{manifest_id}' endpoints
    Manifest *NotificationsWebhooksEncodingManifestAPI

}

// NewNotificationsWebhooksEncodingAPI constructor for NotificationsWebhooksEncodingAPI that takes options as argument
func NewNotificationsWebhooksEncodingAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewNotificationsWebhooksEncodingAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingAPIWithClient constructor for NotificationsWebhooksEncodingAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingAPI {
    a := &NotificationsWebhooksEncodingAPI{apiClient: apiClient}
    a.Encodings = NewNotificationsWebhooksEncodingEncodingsAPIWithClient(apiClient)
    a.Manifest = NewNotificationsWebhooksEncodingManifestAPIWithClient(apiClient)

    return a
}


