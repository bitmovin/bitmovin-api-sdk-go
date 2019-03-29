package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsSpritesCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsSpritesCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsSpritesCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsSpritesCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsSpritesCustomdataApi) GetCustomData(encodingId string, streamId string, spriteId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["sprite_id"] = spriteId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}/customData", &resp, reqParams)
    return resp, err
}
