package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// StreamsConfigAPI intermediary API object with no endpoints
type StreamsConfigAPI struct {
	apiClient *apiclient.APIClient

	// DomainRestriction communicates with '/streams/config/domain-restriction' endpoints
	DomainRestriction *StreamsConfigDomainRestrictionAPI
}

// NewStreamsConfigAPI constructor for StreamsConfigAPI that takes options as argument
func NewStreamsConfigAPI(options ...apiclient.APIClientOption) (*StreamsConfigAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsConfigAPIWithClient(apiClient), nil
}

// NewStreamsConfigAPIWithClient constructor for StreamsConfigAPI that takes an APIClient as argument
func NewStreamsConfigAPIWithClient(apiClient *apiclient.APIClient) *StreamsConfigAPI {
	a := &StreamsConfigAPI{apiClient: apiClient}
	a.DomainRestriction = NewStreamsConfigDomainRestrictionAPIWithClient(apiClient)

	return a
}
