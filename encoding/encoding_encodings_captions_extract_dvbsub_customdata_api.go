package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsCaptionsExtractDvbsubCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsCaptionsExtractDvbsubCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsExtractDvbsubCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsExtractDvbsubCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsExtractDvbsubCustomdataApi) GetCustomData(encodingId string, subtitleId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["subtitle_id"] = subtitleId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/dvbsub/{subtitle_id}/customData", &resp, reqParams)
    return resp, err
}
