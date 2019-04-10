package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsAkamaiMslCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsAkamaiMslCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsAkamaiMslCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsAkamaiMslCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsAkamaiMslCustomdataApi) GetCustomData(outputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/akamai-msl/{output_id}/customData", &resp, reqParams)
    return resp, err
}
