package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AgentsAssistantAPI intermediary API object with no endpoints
type AgentsAssistantAPI struct {
	apiClient *apiclient.APIClient

	// Sessions communicates with '/agents/assistant/sessions' endpoints
	Sessions *AgentsAssistantSessionsAPI
}

// NewAgentsAssistantAPI constructor for AgentsAssistantAPI that takes options as argument
func NewAgentsAssistantAPI(options ...apiclient.APIClientOption) (*AgentsAssistantAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAgentsAssistantAPIWithClient(apiClient), nil
}

// NewAgentsAssistantAPIWithClient constructor for AgentsAssistantAPI that takes an APIClient as argument
func NewAgentsAssistantAPIWithClient(apiClient *apiclient.APIClient) *AgentsAssistantAPI {
	a := &AgentsAssistantAPI{apiClient: apiClient}
	a.Sessions = NewAgentsAssistantSessionsAPIWithClient(apiClient)

	return a
}
