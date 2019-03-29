package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsTypeApi(configs ...func(*common.ApiClient)) (*EncodingManifestsTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsTypeApi) Get(manifestId string) (*model.ManifestTypeResponse, error) {
    var resp *model.ManifestTypeResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/{manifest_id}/type", &resp, reqParams)
    return resp, err
}
