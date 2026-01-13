package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsLiveEsamAPI intermediary API object with no endpoints
type EncodingEncodingsLiveEsamAPI struct {
	apiClient *apiclient.APIClient

	// MediaPoints communicates with '/encoding/encodings/{encoding_id}/live/esam/media-points' endpoints
	MediaPoints *EncodingEncodingsLiveEsamMediaPointsAPI
}

// NewEncodingEncodingsLiveEsamAPI constructor for EncodingEncodingsLiveEsamAPI that takes options as argument
func NewEncodingEncodingsLiveEsamAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveEsamAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveEsamAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveEsamAPIWithClient constructor for EncodingEncodingsLiveEsamAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveEsamAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveEsamAPI {
	a := &EncodingEncodingsLiveEsamAPI{apiClient: apiClient}
	a.MediaPoints = NewEncodingEncodingsLiveEsamMediaPointsAPIWithClient(apiClient)

	return a
}
