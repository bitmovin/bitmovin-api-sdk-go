package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInfrastructureKubernetesConfigurationAPI communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration' endpoints
type EncodingInfrastructureKubernetesConfigurationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureKubernetesConfigurationAPI constructor for EncodingInfrastructureKubernetesConfigurationAPI that takes options as argument
func NewEncodingInfrastructureKubernetesConfigurationAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureKubernetesConfigurationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureKubernetesConfigurationAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureKubernetesConfigurationAPIWithClient constructor for EncodingInfrastructureKubernetesConfigurationAPI that takes an APIClient as argument
func NewEncodingInfrastructureKubernetesConfigurationAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureKubernetesConfigurationAPI {
	a := &EncodingInfrastructureKubernetesConfigurationAPI{apiClient: apiClient}
	return a
}

// Get Retrieve Kubernetes Cluster Configuration
func (api *EncodingInfrastructureKubernetesConfigurationAPI) Get(infrastructureId string) (*model.KubernetesClusterConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.KubernetesClusterConfiguration
	err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update Kubernetes Cluster Configuration
func (api *EncodingInfrastructureKubernetesConfigurationAPI) Update(infrastructureId string, kubernetesClusterConfiguration model.KubernetesClusterConfiguration) (*model.KubernetesClusterConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.KubernetesClusterConfiguration
	err := api.apiClient.Put("/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration", &kubernetesClusterConfiguration, &responseModel, reqParams)
	return &responseModel, err
}
