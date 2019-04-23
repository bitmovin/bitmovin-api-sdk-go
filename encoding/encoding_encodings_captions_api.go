package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsCaptionsApi struct {
    apiClient *common.ApiClient
    Extract *EncodingEncodingsCaptionsExtractApi
    Scc *EncodingEncodingsCaptionsSccApi
}

func NewEncodingEncodingsCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsApi{apiClient: apiClient}

    extractApi, err := NewEncodingEncodingsCaptionsExtractApi(configs...)
    api.Extract = extractApi
    sccApi, err := NewEncodingEncodingsCaptionsSccApi(configs...)
    api.Scc = sccApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

