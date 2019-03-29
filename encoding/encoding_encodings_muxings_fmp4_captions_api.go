package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsMuxingsFmp4CaptionsApi struct {
    apiClient *common.ApiClient
    Webvtt *EncodingEncodingsMuxingsFmp4CaptionsWebvttApi
    Ttml *EncodingEncodingsMuxingsFmp4CaptionsTtmlApi
}

func NewEncodingEncodingsMuxingsFmp4CaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4CaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4CaptionsApi{apiClient: apiClient}

    webvttApi, err := NewEncodingEncodingsMuxingsFmp4CaptionsWebvttApi(configs...)
    api.Webvtt = webvttApi
    ttmlApi, err := NewEncodingEncodingsMuxingsFmp4CaptionsTtmlApi(configs...)
    api.Ttml = ttmlApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

