package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingHistoryAPI intermediary API object with no endpoints
type EncodingHistoryAPI struct {
	apiClient *apiclient.APIClient

	// Encodings communicates with '/encoding/history/encodings' endpoints
	Encodings *EncodingHistoryEncodingsAPI
}

// NewEncodingHistoryAPI constructor for EncodingHistoryAPI that takes options as argument
func NewEncodingHistoryAPI(options ...apiclient.APIClientOption) (*EncodingHistoryAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingHistoryAPIWithClient(apiClient), nil
}

// NewEncodingHistoryAPIWithClient constructor for EncodingHistoryAPI that takes an APIClient as argument
func NewEncodingHistoryAPIWithClient(apiClient *apiclient.APIClient) *EncodingHistoryAPI {
	a := &EncodingHistoryAPI{apiClient: apiClient}
	a.Encodings = NewEncodingHistoryEncodingsAPIWithClient(apiClient)

	return a
}
