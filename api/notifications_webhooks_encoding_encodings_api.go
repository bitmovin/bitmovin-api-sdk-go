package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// NotificationsWebhooksEncodingEncodingsAPI intermediary API object with no endpoints
type NotificationsWebhooksEncodingEncodingsAPI struct {
	apiClient *apiclient.APIClient

	// Finished communicates with '/notifications/webhooks/encoding/encodings/finished' endpoints
	Finished *NotificationsWebhooksEncodingEncodingsFinishedAPI
	// Error communicates with '/notifications/webhooks/encoding/encodings/error' endpoints
	Error *NotificationsWebhooksEncodingEncodingsErrorAPI
	// TransferError communicates with '/notifications/webhooks/encoding/encodings/transfer-error' endpoints
	TransferError *NotificationsWebhooksEncodingEncodingsTransferErrorAPI
	// LiveInputStreamChanged communicates with '/notifications/webhooks/encoding/encodings/live-input-stream-changed' endpoints
	LiveInputStreamChanged *NotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPI
	// EncodingStatusChanged communicates with '/notifications/webhooks/encoding/encodings/encoding-status-changed' endpoints
	EncodingStatusChanged *NotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPI
}

// NewNotificationsWebhooksEncodingEncodingsAPI constructor for NotificationsWebhooksEncodingEncodingsAPI that takes options as argument
func NewNotificationsWebhooksEncodingEncodingsAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingEncodingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingEncodingsAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingEncodingsAPIWithClient constructor for NotificationsWebhooksEncodingEncodingsAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingEncodingsAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingEncodingsAPI {
	a := &NotificationsWebhooksEncodingEncodingsAPI{apiClient: apiClient}
	a.Finished = NewNotificationsWebhooksEncodingEncodingsFinishedAPIWithClient(apiClient)
	a.Error = NewNotificationsWebhooksEncodingEncodingsErrorAPIWithClient(apiClient)
	a.TransferError = NewNotificationsWebhooksEncodingEncodingsTransferErrorAPIWithClient(apiClient)
	a.LiveInputStreamChanged = NewNotificationsWebhooksEncodingEncodingsLiveInputStreamChangedAPIWithClient(apiClient)
	a.EncodingStatusChanged = NewNotificationsWebhooksEncodingEncodingsEncodingStatusChangedAPIWithClient(apiClient)

	return a
}
