package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsCaptionsExtractApi struct {
    apiClient *common.ApiClient
    Cea *EncodingEncodingsCaptionsExtractCeaApi
    Dvbsub *EncodingEncodingsCaptionsExtractDvbsubApi
    Ttml *EncodingEncodingsCaptionsExtractTtmlApi
    Webvtt *EncodingEncodingsCaptionsExtractWebvttApi
}

func NewEncodingEncodingsCaptionsExtractApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsExtractApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsExtractApi{apiClient: apiClient}

    ceaApi, err := NewEncodingEncodingsCaptionsExtractCeaApi(configs...)
    api.Cea = ceaApi
    dvbsubApi, err := NewEncodingEncodingsCaptionsExtractDvbsubApi(configs...)
    api.Dvbsub = dvbsubApi
    ttmlApi, err := NewEncodingEncodingsCaptionsExtractTtmlApi(configs...)
    api.Ttml = ttmlApi
    webvttApi, err := NewEncodingEncodingsCaptionsExtractWebvttApi(configs...)
    api.Webvtt = webvttApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

