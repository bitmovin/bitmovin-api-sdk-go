package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AgentsAPI intermediary API object with no endpoints
type AgentsAPI struct {
	apiClient *apiclient.APIClient

	// Aisa intermediary API object with no endpoints
	Aisa *AgentsAisaAPI
}

// NewAgentsAPI constructor for AgentsAPI that takes options as argument
func NewAgentsAPI(options ...apiclient.APIClientOption) (*AgentsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAgentsAPIWithClient(apiClient), nil
}

// NewAgentsAPIWithClient constructor for AgentsAPI that takes an APIClient as argument
func NewAgentsAPIWithClient(apiClient *apiclient.APIClient) *AgentsAPI {
	a := &AgentsAPI{apiClient: apiClient}
	a.Aisa = NewAgentsAisaAPIWithClient(apiClient)

	return a
}
