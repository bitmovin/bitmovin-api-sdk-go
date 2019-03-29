package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsDashCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashCustomdataApi) GetCustomData(manifestId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/customData", &resp, reqParams)
    return resp, err
}
