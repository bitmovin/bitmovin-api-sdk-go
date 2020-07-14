package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioDtsPassthroughCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioDtsPassthroughCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioDtsPassthroughCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioDtsPassthroughCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioDtsPassthroughCustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/dts-passthrough/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

