package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioAc3CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioAc3CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioAc3CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioAc3CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioAc3CustomdataApi) GetCustomData(configurationId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/audio/ac3/{configuration_id}/customData", &resp, reqParams)
    return resp, err
}
