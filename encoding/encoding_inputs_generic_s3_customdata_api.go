package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsGenericS3CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsGenericS3CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsGenericS3CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGenericS3CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGenericS3CustomdataApi) GetCustomData(inputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
	}
    err := api.apiClient.Get("/encoding/inputs/generic-s3/{input_id}/customData", &resp, reqParams)
    return resp, err
}
