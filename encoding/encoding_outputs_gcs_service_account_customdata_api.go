package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsGcsServiceAccountCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsGcsServiceAccountCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsGcsServiceAccountCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsGcsServiceAccountCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsGcsServiceAccountCustomdataApi) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/outputs/gcs-service-account/{output_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

