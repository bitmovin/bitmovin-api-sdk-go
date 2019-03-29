package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureKubernetesPrewarmedDeploymentApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesPrewarmedDeploymentApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesPrewarmedDeploymentApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesPrewarmedDeploymentApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) List(infrastructureId string, queryParams ...func(*query.PrewarmEncoderSettingsListQueryParams)) (*pagination.PrewarmEncoderSettingssListPagination, error) {
    queryParameters := &query.PrewarmEncoderSettingsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.PrewarmEncoderSettingssListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment", &resp, reqParams)
    return resp, err
}
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) Create(infrastructureId string, prewarmEncoderSettings model.PrewarmEncoderSettings) (*model.PrewarmEncoderSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }
    payload := model.PrewarmEncoderSettings(prewarmEncoderSettings)
    
    err := api.apiClient.Post("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment", &payload, reqParams)
    return &payload, err
}
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) Delete(infrastructureId string, deploymentId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["deployment_id"] = deploymentId
	}
    err := api.apiClient.Delete("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment/{deployment_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) Get(infrastructureId string, deploymentId string) (*model.PrewarmEncoderSettings, error) {
    var resp *model.PrewarmEncoderSettings
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["deployment_id"] = deploymentId
	}
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment/{deployment_id}", &resp, reqParams)
    return resp, err
}
