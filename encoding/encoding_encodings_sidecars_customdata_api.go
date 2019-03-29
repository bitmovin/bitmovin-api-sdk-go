package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsSidecarsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsSidecarsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSidecarsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSidecarsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsSidecarsCustomdataApi) GetCustomData(encodingId string, sidecarId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}/customData", &resp, reqParams)
    return resp, err
}
