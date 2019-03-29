package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4DrmApi struct {
    apiClient *common.ApiClient
    Playready *EncodingEncodingsMuxingsMp4DrmPlayreadyApi
    Clearkey *EncodingEncodingsMuxingsMp4DrmClearkeyApi
    Widevine *EncodingEncodingsMuxingsMp4DrmWidevineApi
    Marlin *EncodingEncodingsMuxingsMp4DrmMarlinApi
    Cenc *EncodingEncodingsMuxingsMp4DrmCencApi
}

func NewEncodingEncodingsMuxingsMp4DrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmApi{apiClient: apiClient}

    playreadyApi, err := NewEncodingEncodingsMuxingsMp4DrmPlayreadyApi(configs...)
    api.Playready = playreadyApi
    clearkeyApi, err := NewEncodingEncodingsMuxingsMp4DrmClearkeyApi(configs...)
    api.Clearkey = clearkeyApi
    widevineApi, err := NewEncodingEncodingsMuxingsMp4DrmWidevineApi(configs...)
    api.Widevine = widevineApi
    marlinApi, err := NewEncodingEncodingsMuxingsMp4DrmMarlinApi(configs...)
    api.Marlin = marlinApi
    cencApi, err := NewEncodingEncodingsMuxingsMp4DrmCencApi(configs...)
    api.Cenc = cencApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    var resp *pagination.DrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm", &resp, reqParams)
    return resp, err
}
