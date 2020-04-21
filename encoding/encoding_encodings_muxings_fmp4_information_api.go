package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsFmp4InformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsFmp4InformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4InformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4InformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4InformationApi) Get(encodingId string, muxingId string) (*model.Fmp4MuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.Fmp4MuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

