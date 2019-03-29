package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsStreamsSubtitlesApi struct {
    apiClient *common.ApiClient
    Dvbsub *EncodingEncodingsStreamsSubtitlesDvbsubApi
}

func NewEncodingEncodingsStreamsSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsSubtitlesApi{apiClient: apiClient}

    dvbsubApi, err := NewEncodingEncodingsStreamsSubtitlesDvbsubApi(configs...)
    api.Dvbsub = dvbsubApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

