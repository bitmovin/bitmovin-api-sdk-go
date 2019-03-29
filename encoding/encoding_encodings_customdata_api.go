package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCustomdataApi) GetCustomData(encodingId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/customData", &resp, reqParams)
    return resp, err
}
