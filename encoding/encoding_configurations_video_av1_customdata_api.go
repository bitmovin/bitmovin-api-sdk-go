package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsVideoAv1CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsVideoAv1CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoAv1CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoAv1CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoAv1CustomdataApi) GetCustomData(configurationId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/video/av1/{configuration_id}/customData", &resp, reqParams)
    return resp, err
}
