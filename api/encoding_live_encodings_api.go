package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingLiveEncodingsAPI intermediary API object with no endpoints
type EncodingLiveEncodingsAPI struct {
	apiClient *apiclient.APIClient

	// Actions communicates with '/encoding/live/encodings/{encoding_id}/actions/update-rtmp-ingest-points' endpoints
	Actions *EncodingLiveEncodingsActionsAPI
	// DnsMappings communicates with '/encoding/live/encodings/{encoding_id}/dns-mappings' endpoints
	DnsMappings *EncodingLiveEncodingsDnsMappingsAPI
}

// NewEncodingLiveEncodingsAPI constructor for EncodingLiveEncodingsAPI that takes options as argument
func NewEncodingLiveEncodingsAPI(options ...apiclient.APIClientOption) (*EncodingLiveEncodingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveEncodingsAPIWithClient(apiClient), nil
}

// NewEncodingLiveEncodingsAPIWithClient constructor for EncodingLiveEncodingsAPI that takes an APIClient as argument
func NewEncodingLiveEncodingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveEncodingsAPI {
	a := &EncodingLiveEncodingsAPI{apiClient: apiClient}
	a.Actions = NewEncodingLiveEncodingsActionsAPIWithClient(apiClient)
	a.DnsMappings = NewEncodingLiveEncodingsDnsMappingsAPIWithClient(apiClient)

	return a
}
