package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsHlsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsCustomdataApi) GetCustomData(manifestId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/customData", &resp, reqParams)
    return resp, err
}
