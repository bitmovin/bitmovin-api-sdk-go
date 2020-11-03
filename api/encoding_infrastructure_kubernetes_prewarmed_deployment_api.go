package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInfrastructureKubernetesPrewarmedDeploymentAPI communicates with '/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment' endpoints
type EncodingInfrastructureKubernetesPrewarmedDeploymentAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInfrastructureKubernetesPrewarmedDeploymentAPI constructor for EncodingInfrastructureKubernetesPrewarmedDeploymentAPI that takes options as argument
func NewEncodingInfrastructureKubernetesPrewarmedDeploymentAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureKubernetesPrewarmedDeploymentAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInfrastructureKubernetesPrewarmedDeploymentAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureKubernetesPrewarmedDeploymentAPIWithClient constructor for EncodingInfrastructureKubernetesPrewarmedDeploymentAPI that takes an APIClient as argument
func NewEncodingInfrastructureKubernetesPrewarmedDeploymentAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureKubernetesPrewarmedDeploymentAPI {
    a := &EncodingInfrastructureKubernetesPrewarmedDeploymentAPI{apiClient: apiClient}
    return a
}

// Create Prewarm Encoders
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentAPI) Create(infrastructureId string, prewarmEncoderSettings model.PrewarmEncoderSettings) (*model.PrewarmEncoderSettings, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel model.PrewarmEncoderSettings
    err := api.apiClient.Post("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment", &prewarmEncoderSettings, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Prewarmed Encoders
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentAPI) Delete(infrastructureId string, deploymentId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["deployment_id"] = deploymentId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment/{deployment_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Prewarmed Encoders
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentAPI) Get(infrastructureId string, deploymentId string) (*model.PrewarmEncoderSettings, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["deployment_id"] = deploymentId
    }

    var responseModel model.PrewarmEncoderSettings
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment/{deployment_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Prewarmed Encoders
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentAPI) List(infrastructureId string, queryParams ...func(*EncodingInfrastructureKubernetesPrewarmedDeploymentAPIListQueryParams)) (*pagination.PrewarmEncoderSettingssListPagination, error) {
    queryParameters := &EncodingInfrastructureKubernetesPrewarmedDeploymentAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.PrewarmEncoderSettingssListPagination
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInfrastructureKubernetesPrewarmedDeploymentAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInfrastructureKubernetesPrewarmedDeploymentAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingInfrastructureKubernetesPrewarmedDeploymentAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


