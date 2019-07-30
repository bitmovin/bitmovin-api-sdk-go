package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsLiveStopInsertedContentApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveStopInsertedContentApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveStopInsertedContentApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveStopInsertedContentApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveStopInsertedContentApi) Create(encodingId string) (error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/stop-inserted-content", nil, nil, reqParams)
    return err
}

