package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi) GetCustomData(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/marlin/{drm_id}/customData", &resp, reqParams)
    return resp, err
}
