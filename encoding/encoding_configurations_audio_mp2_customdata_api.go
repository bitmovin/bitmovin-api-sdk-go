package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioMp2CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioMp2CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioMp2CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioMp2CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioMp2CustomdataApi) GetCustomData(configurationId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/audio/mp2/{configuration_id}/customData", &resp, reqParams)
    return resp, err
}
