package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInfrastructureKubernetesConfigurationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesConfigurationApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesConfigurationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesConfigurationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesConfigurationApi) Get(infrastructureId string) (*model.KubernetesClusterConfiguration, error) {
    var resp *model.KubernetesClusterConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
	}
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration", &resp, reqParams)
    return resp, err
}
func (api *EncodingInfrastructureKubernetesConfigurationApi) Update(infrastructureId string, kubernetesClusterConfiguration *model.KubernetesClusterConfiguration) (*model.KubernetesClusterConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }
    payload := model.KubernetesClusterConfiguration(*kubernetesClusterConfiguration)
    
    err := api.apiClient.Put("/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration", &payload, reqParams)
    return &payload, err
}
