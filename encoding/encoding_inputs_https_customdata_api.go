package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsHttpsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsHttpsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsHttpsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsHttpsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsHttpsCustomdataApi) GetCustomData(inputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/https/{input_id}/customData", &resp, reqParams)
    return resp, err
}
