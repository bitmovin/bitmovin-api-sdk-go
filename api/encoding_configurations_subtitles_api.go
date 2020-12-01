package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingConfigurationsSubtitlesAPI intermediary API object with no endpoints
type EncodingConfigurationsSubtitlesAPI struct {
	apiClient *apiclient.APIClient

	// Webvtt communicates with '/encoding/configurations/subtitles/webvtt' endpoints
	Webvtt *EncodingConfigurationsSubtitlesWebvttAPI
}

// NewEncodingConfigurationsSubtitlesAPI constructor for EncodingConfigurationsSubtitlesAPI that takes options as argument
func NewEncodingConfigurationsSubtitlesAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsSubtitlesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsSubtitlesAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsSubtitlesAPIWithClient constructor for EncodingConfigurationsSubtitlesAPI that takes an APIClient as argument
func NewEncodingConfigurationsSubtitlesAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsSubtitlesAPI {
	a := &EncodingConfigurationsSubtitlesAPI{apiClient: apiClient}
	a.Webvtt = NewEncodingConfigurationsSubtitlesWebvttAPIWithClient(apiClient)

	return a
}
