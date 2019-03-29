package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsInputStreamsTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsTypeApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsTypeApi) Get(encodingId string, inputStreamId string) (*model.InputStreamTypeResponse, error) {
    var resp *model.InputStreamTypeResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}/type", &resp, reqParams)
    return resp, err
}
