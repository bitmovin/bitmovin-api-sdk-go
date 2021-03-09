package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsMuxingsPackedAudioDrmAPI intermediary API object with no endpoints
type EncodingEncodingsMuxingsPackedAudioDrmAPI struct {
	apiClient *apiclient.APIClient

	// Aes communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes' endpoints
	Aes *EncodingEncodingsMuxingsPackedAudioDrmAesAPI
}

// NewEncodingEncodingsMuxingsPackedAudioDrmAPI constructor for EncodingEncodingsMuxingsPackedAudioDrmAPI that takes options as argument
func NewEncodingEncodingsMuxingsPackedAudioDrmAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsPackedAudioDrmAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsPackedAudioDrmAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsPackedAudioDrmAPIWithClient constructor for EncodingEncodingsMuxingsPackedAudioDrmAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsPackedAudioDrmAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsPackedAudioDrmAPI {
	a := &EncodingEncodingsMuxingsPackedAudioDrmAPI{apiClient: apiClient}
	a.Aes = NewEncodingEncodingsMuxingsPackedAudioDrmAesAPIWithClient(apiClient)

	return a
}
