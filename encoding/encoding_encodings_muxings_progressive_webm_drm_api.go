package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
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

