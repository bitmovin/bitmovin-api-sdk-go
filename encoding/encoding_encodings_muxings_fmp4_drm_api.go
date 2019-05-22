package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4DrmApi struct {
    apiClient *common.ApiClient
    Widevine *EncodingEncodingsMuxingsFmp4DrmWidevineApi
    Playready *EncodingEncodingsMuxingsFmp4DrmPlayreadyApi
    Primetime *EncodingEncodingsMuxingsFmp4DrmPrimetimeApi
    Fairplay *EncodingEncodingsMuxingsFmp4DrmFairplayApi
    Marlin *EncodingEncodingsMuxingsFmp4DrmMarlinApi
    Clearkey *EncodingEncodingsMuxingsFmp4DrmClearkeyApi
    Cenc *EncodingEncodingsMuxingsFmp4DrmCencApi
}

func NewEncodingEncodingsMuxingsFmp4DrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4DrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4DrmApi{apiClient: apiClient}

    widevineApi, err := NewEncodingEncodingsMuxingsFmp4DrmWidevineApi(configs...)
    api.Widevine = widevineApi
    playreadyApi, err := NewEncodingEncodingsMuxingsFmp4DrmPlayreadyApi(configs...)
    api.Playready = playreadyApi
    primetimeApi, err := NewEncodingEncodingsMuxingsFmp4DrmPrimetimeApi(configs...)
    api.Primetime = primetimeApi
    fairplayApi, err := NewEncodingEncodingsMuxingsFmp4DrmFairplayApi(configs...)
    api.Fairplay = fairplayApi
    marlinApi, err := NewEncodingEncodingsMuxingsFmp4DrmMarlinApi(configs...)
    api.Marlin = marlinApi
    clearkeyApi, err := NewEncodingEncodingsMuxingsFmp4DrmClearkeyApi(configs...)
    api.Clearkey = clearkeyApi
    cencApi, err := NewEncodingEncodingsMuxingsFmp4DrmCencApi(configs...)
    api.Cenc = cencApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4DrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm", &responseModel, reqParams)
    return responseModel, err
}

