package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsGcsServiceAccountCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsGcsServiceAccountCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsGcsServiceAccountCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGcsServiceAccountCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGcsServiceAccountCustomdataApi) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/inputs/gcs-service-account/{input_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

