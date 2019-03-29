package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4CaptionsWebvttCustomdataApi) GetCustomData(encodingId string, muxingId string, captionsId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/webvtt/{captions_id}/customData", &resp, reqParams)
    return resp, err
}
