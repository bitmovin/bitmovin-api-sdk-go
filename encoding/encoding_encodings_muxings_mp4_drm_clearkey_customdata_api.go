package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataApi) GetCustomData(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey/{drm_id}/customData", &resp, reqParams)
    return resp, err
}
