package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// NotificationsWebhooksAPI intermediary API object with no endpoints
type NotificationsWebhooksAPI struct {
	apiClient *apiclient.APIClient

	// Encoding intermediary API object with no endpoints
	Encoding *NotificationsWebhooksEncodingAPI
}

// NewNotificationsWebhooksAPI constructor for NotificationsWebhooksAPI that takes options as argument
func NewNotificationsWebhooksAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksAPIWithClient constructor for NotificationsWebhooksAPI that takes an APIClient as argument
func NewNotificationsWebhooksAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksAPI {
	a := &NotificationsWebhooksAPI{apiClient: apiClient}
	a.Encoding = NewNotificationsWebhooksEncodingAPIWithClient(apiClient)

	return a
}
