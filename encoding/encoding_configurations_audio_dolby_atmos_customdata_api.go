package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioDolbyAtmosCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioDolbyAtmosCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioDolbyAtmosCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioDolbyAtmosCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioDolbyAtmosCustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/dolby-atmos/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

