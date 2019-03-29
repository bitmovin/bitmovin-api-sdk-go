package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsSubtitlesApi struct {
    apiClient *common.ApiClient
    Dvbsub *EncodingEncodingsSubtitlesDvbsubApi
}

func NewEncodingEncodingsSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSubtitlesApi{apiClient: apiClient}

    dvbsubApi, err := NewEncodingEncodingsSubtitlesDvbsubApi(configs...)
    api.Dvbsub = dvbsubApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

