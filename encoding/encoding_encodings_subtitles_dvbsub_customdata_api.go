package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsSubtitlesDvbsubCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsSubtitlesDvbsubCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSubtitlesDvbsubCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSubtitlesDvbsubCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsSubtitlesDvbsubCustomdataApi) GetCustomData(encodingId string, subtitleId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/subtitles/dvbsub/{subtitle_id}/customData", &resp, reqParams)
    return resp, err
}
