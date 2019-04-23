package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsCaptionsExtractTtmlCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsCaptionsExtractTtmlCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsExtractTtmlCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsExtractTtmlCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsExtractTtmlCustomdataApi) GetCustomData(encodingId string, captionsId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/ttml/{captions_id}/customData", &resp, reqParams)
    return resp, err
}
