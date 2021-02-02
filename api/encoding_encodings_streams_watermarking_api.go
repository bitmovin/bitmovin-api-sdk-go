package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsStreamsWatermarkingAPI intermediary API object with no endpoints
type EncodingEncodingsStreamsWatermarkingAPI struct {
	apiClient *apiclient.APIClient

	// NexguardFileMarker communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker' endpoints
	NexguardFileMarker *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI
}

// NewEncodingEncodingsStreamsWatermarkingAPI constructor for EncodingEncodingsStreamsWatermarkingAPI that takes options as argument
func NewEncodingEncodingsStreamsWatermarkingAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsWatermarkingAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsWatermarkingAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsWatermarkingAPIWithClient constructor for EncodingEncodingsStreamsWatermarkingAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsWatermarkingAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsWatermarkingAPI {
	a := &EncodingEncodingsStreamsWatermarkingAPI{apiClient: apiClient}
	a.NexguardFileMarker = NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIWithClient(apiClient)

	return a
}
