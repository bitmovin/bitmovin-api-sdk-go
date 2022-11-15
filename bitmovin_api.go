package bitmovin

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/api"
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// BitmovinAPI intermediary API object with no endpoints
type BitmovinAPI struct {
	apiClient *apiclient.APIClient

	// Account intermediary API object with no endpoints
	Account *api.AccountAPI
	// Analytics intermediary API object with no endpoints
	Analytics *api.AnalyticsAPI
	// Encoding intermediary API object with no endpoints
	Encoding *api.EncodingAPI
	// General intermediary API object with no endpoints
	General *api.GeneralAPI
	// Notifications communicates with '/notifications' endpoints
	Notifications *api.NotificationsAPI
	// Player intermediary API object with no endpoints
	Player *api.PlayerAPI
	// Streams intermediary API object with no endpoints
	Streams *api.StreamsAPI
}

// NewBitmovinAPI constructor for BitmovinAPI that takes options as argument
func NewBitmovinAPI(options ...apiclient.APIClientOption) (*BitmovinAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewBitmovinAPIWithClient(apiClient), nil
}

// NewBitmovinAPIWithClient constructor for BitmovinAPI that takes an APIClient as argument
func NewBitmovinAPIWithClient(apiClient *apiclient.APIClient) *BitmovinAPI {
	a := &BitmovinAPI{apiClient: apiClient}
	a.Account = api.NewAccountAPIWithClient(apiClient)
	a.Analytics = api.NewAnalyticsAPIWithClient(apiClient)
	a.Encoding = api.NewEncodingAPIWithClient(apiClient)
	a.General = api.NewGeneralAPIWithClient(apiClient)
	a.Notifications = api.NewNotificationsAPIWithClient(apiClient)
	a.Player = api.NewPlayerAPIWithClient(apiClient)
	a.Streams = api.NewStreamsAPIWithClient(apiClient)

	return a
}
