package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsAkamaiNetstorageCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsAkamaiNetstorageCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsAkamaiNetstorageCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsAkamaiNetstorageCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsAkamaiNetstorageCustomdataApi) GetCustomData(outputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/akamai-netstorage/{output_id}/customData", &resp, reqParams)
    return resp, err
}
