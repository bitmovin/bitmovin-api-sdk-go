package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafDrmApi struct {
    apiClient *common.ApiClient
    Widevine *EncodingEncodingsMuxingsCmafDrmWidevineApi
    Playready *EncodingEncodingsMuxingsCmafDrmPlayreadyApi
    Primetime *EncodingEncodingsMuxingsCmafDrmPrimetimeApi
    Fairplay *EncodingEncodingsMuxingsCmafDrmFairplayApi
    Marlin *EncodingEncodingsMuxingsCmafDrmMarlinApi
    Clearkey *EncodingEncodingsMuxingsCmafDrmClearkeyApi
    Cenc *EncodingEncodingsMuxingsCmafDrmCencApi
}

func NewEncodingEncodingsMuxingsCmafDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmApi{apiClient: apiClient}

    widevineApi, err := NewEncodingEncodingsMuxingsCmafDrmWidevineApi(configs...)
    api.Widevine = widevineApi
    playreadyApi, err := NewEncodingEncodingsMuxingsCmafDrmPlayreadyApi(configs...)
    api.Playready = playreadyApi
    primetimeApi, err := NewEncodingEncodingsMuxingsCmafDrmPrimetimeApi(configs...)
    api.Primetime = primetimeApi
    fairplayApi, err := NewEncodingEncodingsMuxingsCmafDrmFairplayApi(configs...)
    api.Fairplay = fairplayApi
    marlinApi, err := NewEncodingEncodingsMuxingsCmafDrmMarlinApi(configs...)
    api.Marlin = marlinApi
    clearkeyApi, err := NewEncodingEncodingsMuxingsCmafDrmClearkeyApi(configs...)
    api.Clearkey = clearkeyApi
    cencApi, err := NewEncodingEncodingsMuxingsCmafDrmCencApi(configs...)
    api.Cenc = cencApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    var resp *pagination.DrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm", &resp, reqParams)
    return resp, err
}
