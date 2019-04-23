package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubCustomdataApi) GetCustomData(encodingId string, streamId string, subtitleId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}/customData", &resp, reqParams)
    return resp, err
}
