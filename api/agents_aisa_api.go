package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AgentsAisaAPI intermediary API object with no endpoints
type AgentsAisaAPI struct {
	apiClient *apiclient.APIClient

	// Sessions communicates with '/agents/aisa/sessions' endpoints
	Sessions *AgentsAisaSessionsAPI
}

// NewAgentsAisaAPI constructor for AgentsAisaAPI that takes options as argument
func NewAgentsAisaAPI(options ...apiclient.APIClientOption) (*AgentsAisaAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAgentsAisaAPIWithClient(apiClient), nil
}

// NewAgentsAisaAPIWithClient constructor for AgentsAisaAPI that takes an APIClient as argument
func NewAgentsAisaAPIWithClient(apiClient *apiclient.APIClient) *AgentsAisaAPI {
	a := &AgentsAisaAPI{apiClient: apiClient}
	a.Sessions = NewAgentsAisaSessionsAPIWithClient(apiClient)

	return a
}
