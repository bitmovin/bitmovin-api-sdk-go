package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsCaptionsCeaSccCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsCaptionsCeaSccCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsCaptionsCeaSccCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsCaptionsCeaSccCustomdataApi) GetCustomData(encodingId string, streamId string, captionsId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}/customData", &resp, reqParams)
    return resp, err
}
