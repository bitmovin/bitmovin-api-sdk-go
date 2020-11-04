package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingInfrastructureKubernetesStatusAPI communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/status' endpoints
type EncodingInfrastructureKubernetesStatusAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureKubernetesStatusAPI constructor for EncodingInfrastructureKubernetesStatusAPI that takes options as argument
func NewEncodingInfrastructureKubernetesStatusAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureKubernetesStatusAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureKubernetesStatusAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureKubernetesStatusAPIWithClient constructor for EncodingInfrastructureKubernetesStatusAPI that takes an APIClient as argument
func NewEncodingInfrastructureKubernetesStatusAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureKubernetesStatusAPI {
	a := &EncodingInfrastructureKubernetesStatusAPI{apiClient: apiClient}
	return a
}

// Get Kubernetes Cluster Status
func (api *EncodingInfrastructureKubernetesStatusAPI) Get(infrastructureId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/status", nil, nil, reqParams)
	return err
}
