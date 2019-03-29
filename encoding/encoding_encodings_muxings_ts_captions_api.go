package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsMuxingsTsCaptionsApi struct {
    apiClient *common.ApiClient
    Cea *EncodingEncodingsMuxingsTsCaptionsCeaApi
}

func NewEncodingEncodingsMuxingsTsCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsCaptionsApi{apiClient: apiClient}

    ceaApi, err := NewEncodingEncodingsMuxingsTsCaptionsCeaApi(configs...)
    api.Cea = ceaApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

