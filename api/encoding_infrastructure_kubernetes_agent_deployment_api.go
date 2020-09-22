package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingInfrastructureKubernetesAgentDeploymentAPI communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/agent-deployment' endpoints
type EncodingInfrastructureKubernetesAgentDeploymentAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureKubernetesAgentDeploymentAPI constructor for EncodingInfrastructureKubernetesAgentDeploymentAPI that takes options as argument
func NewEncodingInfrastructureKubernetesAgentDeploymentAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureKubernetesAgentDeploymentAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureKubernetesAgentDeploymentAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureKubernetesAgentDeploymentAPIWithClient constructor for EncodingInfrastructureKubernetesAgentDeploymentAPI that takes an APIClient as argument
func NewEncodingInfrastructureKubernetesAgentDeploymentAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureKubernetesAgentDeploymentAPI {
	a := &EncodingInfrastructureKubernetesAgentDeploymentAPI{apiClient: apiClient}
	return a
}

// Get Download bitmovin-agent deployment
func (api *EncodingInfrastructureKubernetesAgentDeploymentAPI) Get(infrastructureId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/agent-deployment", nil, nil, reqParams)
	return err
}
