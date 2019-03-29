package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsS3RoleBasedCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsS3RoleBasedCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsS3RoleBasedCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsS3RoleBasedCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsS3RoleBasedCustomdataApi) GetCustomData(outputId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/encoding/outputs/s3-role-based/{output_id}/customData", &resp, reqParams)
    return resp, err
}
