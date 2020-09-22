package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureKubernetesAPI communicates with '/encoding/infrastructure/kubernetes' endpoints
type EncodingInfrastructureKubernetesAPI struct {
	apiClient *apiclient.APIClient

	// Status communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/status' endpoints
	Status *EncodingInfrastructureKubernetesStatusAPI
	// Customdata communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/customData' endpoints
	Customdata *EncodingInfrastructureKubernetesCustomdataAPI
	// Configuration communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration' endpoints
	Configuration *EncodingInfrastructureKubernetesConfigurationAPI
	// AgentDeployment communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/agent-deployment' endpoints
	AgentDeployment *EncodingInfrastructureKubernetesAgentDeploymentAPI
	// PrewarmedDeployment communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment' endpoints
	PrewarmedDeployment *EncodingInfrastructureKubernetesPrewarmedDeploymentAPI
}

// NewEncodingInfrastructureKubernetesAPI constructor for EncodingInfrastructureKubernetesAPI that takes options as argument
func NewEncodingInfrastructureKubernetesAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureKubernetesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureKubernetesAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureKubernetesAPIWithClient constructor for EncodingInfrastructureKubernetesAPI that takes an APIClient as argument
func NewEncodingInfrastructureKubernetesAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureKubernetesAPI {
	a := &EncodingInfrastructureKubernetesAPI{apiClient: apiClient}
	a.Status = NewEncodingInfrastructureKubernetesStatusAPIWithClient(apiClient)
	a.Customdata = NewEncodingInfrastructureKubernetesCustomdataAPIWithClient(apiClient)
	a.Configuration = NewEncodingInfrastructureKubernetesConfigurationAPIWithClient(apiClient)
	a.AgentDeployment = NewEncodingInfrastructureKubernetesAgentDeploymentAPIWithClient(apiClient)
	a.PrewarmedDeployment = NewEncodingInfrastructureKubernetesPrewarmedDeploymentAPIWithClient(apiClient)

	return a
}

// Create Connect Kubernetes Cluster
func (api *EncodingInfrastructureKubernetesAPI) Create(kubernetesCluster model.KubernetesCluster) (*model.KubernetesCluster, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.KubernetesCluster
	err := api.apiClient.Post("/encoding/infrastructure/kubernetes", &kubernetesCluster, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Disconnect Kubernetes Cluster
func (api *EncodingInfrastructureKubernetesAPI) Delete(infrastructureId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/infrastructure/kubernetes/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Kubernetes Cluster Details
func (api *EncodingInfrastructureKubernetesAPI) Get(infrastructureId string) (*model.KubernetesCluster, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["infrastructure_id"] = infrastructureId
	}

	var responseModel model.KubernetesCluster
	err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Kubernetes Cluster
func (api *EncodingInfrastructureKubernetesAPI) List(queryParams ...func(*EncodingInfrastructureKubernetesAPIListQueryParams)) (*pagination.KubernetesClustersListPagination, error) {
	queryParameters := &EncodingInfrastructureKubernetesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.KubernetesClustersListPagination
	err := api.apiClient.Get("/encoding/infrastructure/kubernetes", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInfrastructureKubernetesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureKubernetesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureKubernetesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
