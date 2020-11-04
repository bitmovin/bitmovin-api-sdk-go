package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsStreamsQcAPI intermediary API object with no endpoints
type EncodingEncodingsStreamsQcAPI struct {
	apiClient *apiclient.APIClient

	// Psnr communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/qc/psnr' endpoints
	Psnr *EncodingEncodingsStreamsQcPsnrAPI
}

// NewEncodingEncodingsStreamsQcAPI constructor for EncodingEncodingsStreamsQcAPI that takes options as argument
func NewEncodingEncodingsStreamsQcAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsQcAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsQcAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsQcAPIWithClient constructor for EncodingEncodingsStreamsQcAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsQcAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsQcAPI {
	a := &EncodingEncodingsStreamsQcAPI{apiClient: apiClient}
	a.Psnr = NewEncodingEncodingsStreamsQcPsnrAPIWithClient(apiClient)

	return a
}
