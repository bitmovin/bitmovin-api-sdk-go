package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveWebmDrmApi struct {
    apiClient *common.ApiClient
    Cenc *EncodingEncodingsMuxingsProgressiveWebmDrmCencApi
    Speke *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeApi
}

func NewEncodingEncodingsMuxingsProgressiveWebmDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveWebmDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveWebmDrmApi{apiClient: apiClient}

    cencApi, err := NewEncodingEncodingsMuxingsProgressiveWebmDrmCencApi(configs...)
    api.Cenc = cencApi
    spekeApi, err := NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeApi(configs...)
    api.Speke = spekeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveWebmDrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm", nil, &responseModel, reqParams)
    return responseModel, err
}

