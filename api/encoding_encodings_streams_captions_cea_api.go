package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsStreamsCaptionsCeaAPI intermediary API object with no endpoints
type EncodingEncodingsStreamsCaptionsCeaAPI struct {
	apiClient *apiclient.APIClient

	// Scc communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc' endpoints
	Scc *EncodingEncodingsStreamsCaptionsCeaSccAPI
	// Srt communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt' endpoints
	Srt *EncodingEncodingsStreamsCaptionsCeaSrtAPI
}

// NewEncodingEncodingsStreamsCaptionsCeaAPI constructor for EncodingEncodingsStreamsCaptionsCeaAPI that takes options as argument
func NewEncodingEncodingsStreamsCaptionsCeaAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCaptionsCeaAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsCaptionsCeaAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCaptionsCeaAPIWithClient constructor for EncodingEncodingsStreamsCaptionsCeaAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCaptionsCeaAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCaptionsCeaAPI {
	a := &EncodingEncodingsStreamsCaptionsCeaAPI{apiClient: apiClient}
	a.Scc = NewEncodingEncodingsStreamsCaptionsCeaSccAPIWithClient(apiClient)
	a.Srt = NewEncodingEncodingsStreamsCaptionsCeaSrtAPIWithClient(apiClient)

	return a
}
