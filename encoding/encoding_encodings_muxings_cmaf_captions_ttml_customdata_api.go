package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi) GetCustomData(encodingId string, muxingId string, captionsId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml/{captions_id}/customData", &resp, reqParams)
    return resp, err
}
