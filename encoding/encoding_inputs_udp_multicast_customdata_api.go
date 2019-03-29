package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsUdpMulticastCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsUdpMulticastCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsUdpMulticastCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsUdpMulticastCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsUdpMulticastCustomdataApi) GetCustomData(inputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/udp-multicast/{input_id}/customData", &resp, reqParams)
    return resp, err
}
