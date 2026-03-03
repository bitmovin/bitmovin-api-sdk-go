package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AgentsAisaSessionsAPI communicates with '/agents/aisa/sessions' endpoints
type AgentsAisaSessionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAgentsAisaSessionsAPI constructor for AgentsAisaSessionsAPI that takes options as argument
func NewAgentsAisaSessionsAPI(options ...apiclient.APIClientOption) (*AgentsAisaSessionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAgentsAisaSessionsAPIWithClient(apiClient), nil
}

// NewAgentsAisaSessionsAPIWithClient constructor for AgentsAisaSessionsAPI that takes an APIClient as argument
func NewAgentsAisaSessionsAPIWithClient(apiClient *apiclient.APIClient) *AgentsAisaSessionsAPI {
	a := &AgentsAisaSessionsAPI{apiClient: apiClient}
	return a
}

// Create Agent Session
func (api *AgentsAisaSessionsAPI) Create(xUserId string) (*model.AgentSessionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AgentSessionResponse
	err := api.apiClient.Post("/agents/aisa/sessions", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Agent Session
func (api *AgentsAisaSessionsAPI) Delete(sessionId string, xUserId string) (*model.AgentSessionResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["session_id"] = sessionId
	}

	var responseModel model.AgentSessionResponse
	err := api.apiClient.Delete("/agents/aisa/sessions/{session_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Agent Session Details
func (api *AgentsAisaSessionsAPI) Get(sessionId string, xUserId string) (*model.AgentSessionHistoryResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["session_id"] = sessionId
	}

	var responseModel model.AgentSessionHistoryResponse
	err := api.apiClient.Get("/agents/aisa/sessions/{session_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Agent Sessions
func (api *AgentsAisaSessionsAPI) List(xUserId string) (*model.AgentSessionListResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AgentSessionListResponse
	err := api.apiClient.Get("/agents/aisa/sessions", nil, &responseModel, reqParams)
	return &responseModel, err
}
