package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInfrastructureKubernetesCustomdataAPI communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/customData' endpoints
type EncodingInfrastructureKubernetesCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureKubernetesCustomdataAPI constructor for EncodingInfrastructureKubernetesCustomdataAPI that takes options as argument
func NewEncodingInfrastructureKubernetesCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureKubernetesCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureKubernetesCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureKubernetesCustomdataAPIWithClient constructor for EncodingInfrastructureKubernetesCustomdataAPI that takes an APIClient as argument
func NewEncodingInfrastructureKubernetesCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureKubernetesCustomdataAPI {
	a := &EncodingInfrastructureKubernetesCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Kubernetes Cluster Custom Data
func (api *EncodingInfrastructureKubernetesCustomdataAPI) Get(infrastructureId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
