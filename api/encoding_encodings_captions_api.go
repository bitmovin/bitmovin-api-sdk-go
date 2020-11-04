package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsCaptionsAPI intermediary API object with no endpoints
type EncodingEncodingsCaptionsAPI struct {
	apiClient *apiclient.APIClient

	// Scc communicates with '/encoding/encodings/{encoding_id}/captions/scc' endpoints
	Scc *EncodingEncodingsCaptionsSccAPI
}

// NewEncodingEncodingsCaptionsAPI constructor for EncodingEncodingsCaptionsAPI that takes options as argument
func NewEncodingEncodingsCaptionsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsCaptionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsCaptionsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsCaptionsAPIWithClient constructor for EncodingEncodingsCaptionsAPI that takes an APIClient as argument
func NewEncodingEncodingsCaptionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsCaptionsAPI {
	a := &EncodingEncodingsCaptionsAPI{apiClient: apiClient}
	a.Scc = NewEncodingEncodingsCaptionsSccAPIWithClient(apiClient)

	return a
}
