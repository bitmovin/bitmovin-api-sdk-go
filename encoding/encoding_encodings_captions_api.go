package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsCaptionsApi struct {
    apiClient *common.ApiClient
    Cea *EncodingEncodingsCaptionsCeaApi
    Webvtt *EncodingEncodingsCaptionsWebvttApi
    Ttml *EncodingEncodingsCaptionsTtmlApi
    Scc *EncodingEncodingsCaptionsSccApi
}

func NewEncodingEncodingsCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsApi{apiClient: apiClient}

    ceaApi, err := NewEncodingEncodingsCaptionsCeaApi(configs...)
    api.Cea = ceaApi
    webvttApi, err := NewEncodingEncodingsCaptionsWebvttApi(configs...)
    api.Webvtt = webvttApi
    ttmlApi, err := NewEncodingEncodingsCaptionsTtmlApi(configs...)
    api.Ttml = ttmlApi
    sccApi, err := NewEncodingEncodingsCaptionsSccApi(configs...)
    api.Scc = sccApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

