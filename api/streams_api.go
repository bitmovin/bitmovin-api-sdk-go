package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// StreamsAPI intermediary API object with no endpoints
type StreamsAPI struct {
	apiClient *apiclient.APIClient

	// Video communicates with '/streams/video' endpoints
	Video *StreamsVideoAPI
	// Live communicates with '/streams/live' endpoints
	Live *StreamsLiveAPI
	// Config intermediary API object with no endpoints
	Config *StreamsConfigAPI
	// Search communicates with '/streams/search' endpoints
	Search *StreamsSearchAPI
	// SigningKeys communicates with '/streams/signing-keys' endpoints
	SigningKeys *StreamsSigningKeysAPI
}

// NewStreamsAPI constructor for StreamsAPI that takes options as argument
func NewStreamsAPI(options ...apiclient.APIClientOption) (*StreamsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsAPIWithClient(apiClient), nil
}

// NewStreamsAPIWithClient constructor for StreamsAPI that takes an APIClient as argument
func NewStreamsAPIWithClient(apiClient *apiclient.APIClient) *StreamsAPI {
	a := &StreamsAPI{apiClient: apiClient}
	a.Video = NewStreamsVideoAPIWithClient(apiClient)
	a.Live = NewStreamsLiveAPIWithClient(apiClient)
	a.Config = NewStreamsConfigAPIWithClient(apiClient)
	a.Search = NewStreamsSearchAPIWithClient(apiClient)
	a.SigningKeys = NewStreamsSigningKeysAPIWithClient(apiClient)

	return a
}
