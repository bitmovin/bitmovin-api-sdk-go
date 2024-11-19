package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingLiveAPI intermediary API object with no endpoints
type EncodingLiveAPI struct {
	apiClient *apiclient.APIClient

	// StreamKeys communicates with '/encoding/live/stream-keys' endpoints
	StreamKeys *EncodingLiveStreamKeysAPI
	// StandbyPools communicates with '/encoding/live/standby-pools' endpoints
	StandbyPools *EncodingLiveStandbyPoolsAPI
}

// NewEncodingLiveAPI constructor for EncodingLiveAPI that takes options as argument
func NewEncodingLiveAPI(options ...apiclient.APIClientOption) (*EncodingLiveAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveAPIWithClient(apiClient), nil
}

// NewEncodingLiveAPIWithClient constructor for EncodingLiveAPI that takes an APIClient as argument
func NewEncodingLiveAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveAPI {
	a := &EncodingLiveAPI{apiClient: apiClient}
	a.StreamKeys = NewEncodingLiveStreamKeysAPIWithClient(apiClient)
	a.StandbyPools = NewEncodingLiveStandbyPoolsAPIWithClient(apiClient)

	return a
}
