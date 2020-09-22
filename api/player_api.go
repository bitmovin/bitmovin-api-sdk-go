package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// PlayerAPI intermediary API object with no endpoints
type PlayerAPI struct {
	apiClient *apiclient.APIClient

	// Channels communicates with '/player/channels' endpoints
	Channels *PlayerChannelsAPI
	// Licenses communicates with '/player/licenses' endpoints
	Licenses *PlayerLicensesAPI
	// CustomBuilds intermediary API object with no endpoints
	CustomBuilds *PlayerCustomBuildsAPI
}

// NewPlayerAPI constructor for PlayerAPI that takes options as argument
func NewPlayerAPI(options ...apiclient.APIClientOption) (*PlayerAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewPlayerAPIWithClient(apiClient), nil
}

// NewPlayerAPIWithClient constructor for PlayerAPI that takes an APIClient as argument
func NewPlayerAPIWithClient(apiClient *apiclient.APIClient) *PlayerAPI {
	a := &PlayerAPI{apiClient: apiClient}
	a.Channels = NewPlayerChannelsAPIWithClient(apiClient)
	a.Licenses = NewPlayerLicensesAPIWithClient(apiClient)
	a.CustomBuilds = NewPlayerCustomBuildsAPIWithClient(apiClient)

	return a
}
