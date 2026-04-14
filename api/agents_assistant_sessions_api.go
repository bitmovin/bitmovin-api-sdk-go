package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AgentsAssistantSessionsAPI communicates with '/agents/assistant/sessions' endpoints
type AgentsAssistantSessionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAgentsAssistantSessionsAPI constructor for AgentsAssistantSessionsAPI that takes options as argument
func NewAgentsAssistantSessionsAPI(options ...apiclient.APIClientOption) (*AgentsAssistantSessionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAgentsAssistantSessionsAPIWithClient(apiClient), nil
}

// NewAgentsAssistantSessionsAPIWithClient constructor for AgentsAssistantSessionsAPI that takes an APIClient as argument
func NewAgentsAssistantSessionsAPIWithClient(apiClient *apiclient.APIClient) *AgentsAssistantSessionsAPI {
	a := &AgentsAssistantSessionsAPI{apiClient: apiClient}
	return a
}

// Create Agent Session
func (api *AgentsAssistantSessionsAPI) Create(xUserId string) (*model.AgentSessionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AgentSessionResponse
	err := api.apiClient.Post("/agents/assistant/sessions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Agent Session
func (api *AgentsAssistantSessionsAPI) Delete(sessionId string, xUserId string) (*model.AgentSessionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["session_id"] = sessionId
	}

	var responseModel model.AgentSessionResponse
	err := api.apiClient.Delete("/agents/assistant/sessions/{session_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Agent Session Details
func (api *AgentsAssistantSessionsAPI) Get(sessionId string, xUserId string) (*model.AgentSessionHistoryResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["session_id"] = sessionId
	}

	var responseModel model.AgentSessionHistoryResponse
	err := api.apiClient.Get("/agents/assistant/sessions/{session_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Agent Sessions
func (api *AgentsAssistantSessionsAPI) List(xUserId string) (*model.AgentSessionListResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AgentSessionListResponse
	err := api.apiClient.Get("/agents/assistant/sessions", nil, &responseModel, reqParams)
	return &responseModel, err
}
