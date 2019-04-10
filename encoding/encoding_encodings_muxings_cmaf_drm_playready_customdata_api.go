package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsCmafDrmPlayreadyCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsCmafDrmPlayreadyCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmPlayreadyCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmPlayreadyCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmPlayreadyCustomdataApi) GetCustomData(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/playready/{drm_id}/customData", &resp, reqParams)
    return resp, err
}
