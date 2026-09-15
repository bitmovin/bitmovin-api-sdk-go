package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// PlayerTestingAPI intermediary API object with no endpoints
type PlayerTestingAPI struct {
	apiClient *apiclient.APIClient

	// CodecCompatibility communicates with '/player/testing/codec-compatibility' endpoints
	CodecCompatibility *PlayerTestingCodecCompatibilityAPI
}

// NewPlayerTestingAPI constructor for PlayerTestingAPI that takes options as argument
func NewPlayerTestingAPI(options ...apiclient.APIClientOption) (*PlayerTestingAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewPlayerTestingAPIWithClient(apiClient), nil
}

// NewPlayerTestingAPIWithClient constructor for PlayerTestingAPI that takes an APIClient as argument
func NewPlayerTestingAPIWithClient(apiClient *apiclient.APIClient) *PlayerTestingAPI {
	a := &PlayerTestingAPI{apiClient: apiClient}
	a.CodecCompatibility = NewPlayerTestingCodecCompatibilityAPIWithClient(apiClient)

	return a
}
