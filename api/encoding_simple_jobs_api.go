package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingSimpleJobsAPI intermediary API object with no endpoints
type EncodingSimpleJobsAPI struct {
	apiClient *apiclient.APIClient

	// Vod communicates with '/encoding/simple/jobs/vod' endpoints
	Vod *EncodingSimpleJobsVodAPI
	// Live communicates with '/encoding/simple/jobs/live' endpoints
	Live *EncodingSimpleJobsLiveAPI
}

// NewEncodingSimpleJobsAPI constructor for EncodingSimpleJobsAPI that takes options as argument
func NewEncodingSimpleJobsAPI(options ...apiclient.APIClientOption) (*EncodingSimpleJobsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingSimpleJobsAPIWithClient(apiClient), nil
}

// NewEncodingSimpleJobsAPIWithClient constructor for EncodingSimpleJobsAPI that takes an APIClient as argument
func NewEncodingSimpleJobsAPIWithClient(apiClient *apiclient.APIClient) *EncodingSimpleJobsAPI {
	a := &EncodingSimpleJobsAPI{apiClient: apiClient}
	a.Vod = NewEncodingSimpleJobsVodAPIWithClient(apiClient)
	a.Live = NewEncodingSimpleJobsLiveAPIWithClient(apiClient)

	return a
}
