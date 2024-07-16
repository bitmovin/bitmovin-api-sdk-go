package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingAPI intermediary API object with no endpoints
type EncodingAPI struct {
	apiClient *apiclient.APIClient

	// Inputs communicates with '/encoding/inputs' endpoints
	Inputs *EncodingInputsAPI
	// Outputs communicates with '/encoding/outputs' endpoints
	Outputs *EncodingOutputsAPI
	// Configurations communicates with '/encoding/configurations' endpoints
	Configurations *EncodingConfigurationsAPI
	// Filters communicates with '/encoding/filters' endpoints
	Filters *EncodingFiltersAPI
	// Encodings communicates with '/encoding/encodings' endpoints
	Encodings *EncodingEncodingsAPI
	// Live intermediary API object with no endpoints
	Live *EncodingLiveAPI
	// Manifests communicates with '/encoding/manifests' endpoints
	Manifests *EncodingManifestsAPI
	// Infrastructure intermediary API object with no endpoints
	Infrastructure *EncodingInfrastructureAPI
	// Statistics communicates with '/encoding/statistics' endpoints
	Statistics *EncodingStatisticsAPI
	// WatchFolders communicates with '/encoding/watch-folders' endpoints
	WatchFolders *EncodingWatchFoldersAPI
	// Simple intermediary API object with no endpoints
	Simple *EncodingSimpleAPI
	// ErrorDefinitions communicates with '/encoding/error-definitions' endpoints
	ErrorDefinitions *EncodingErrorDefinitionsAPI
}

// NewEncodingAPI constructor for EncodingAPI that takes options as argument
func NewEncodingAPI(options ...apiclient.APIClientOption) (*EncodingAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingAPIWithClient(apiClient), nil
}

// NewEncodingAPIWithClient constructor for EncodingAPI that takes an APIClient as argument
func NewEncodingAPIWithClient(apiClient *apiclient.APIClient) *EncodingAPI {
	a := &EncodingAPI{apiClient: apiClient}
	a.Inputs = NewEncodingInputsAPIWithClient(apiClient)
	a.Outputs = NewEncodingOutputsAPIWithClient(apiClient)
	a.Configurations = NewEncodingConfigurationsAPIWithClient(apiClient)
	a.Filters = NewEncodingFiltersAPIWithClient(apiClient)
	a.Encodings = NewEncodingEncodingsAPIWithClient(apiClient)
	a.Live = NewEncodingLiveAPIWithClient(apiClient)
	a.Manifests = NewEncodingManifestsAPIWithClient(apiClient)
	a.Infrastructure = NewEncodingInfrastructureAPIWithClient(apiClient)
	a.Statistics = NewEncodingStatisticsAPIWithClient(apiClient)
	a.WatchFolders = NewEncodingWatchFoldersAPIWithClient(apiClient)
	a.Simple = NewEncodingSimpleAPIWithClient(apiClient)
	a.ErrorDefinitions = NewEncodingErrorDefinitionsAPIWithClient(apiClient)

	return a
}
