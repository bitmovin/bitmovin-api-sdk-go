package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingManifestsHlsMediaApi struct {
    apiClient *common.ApiClient
    CustomTag *EncodingManifestsHlsMediaCustomTagApi
    Type *EncodingManifestsHlsMediaTypeApi
    Video *EncodingManifestsHlsMediaVideoApi
    Audio *EncodingManifestsHlsMediaAudioApi
    Subtitles *EncodingManifestsHlsMediaSubtitlesApi
    Vtt *EncodingManifestsHlsMediaVttApi
    ClosedCaptions *EncodingManifestsHlsMediaClosedCaptionsApi
}

func NewEncodingManifestsHlsMediaApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaApi{apiClient: apiClient}

    customTagApi, err := NewEncodingManifestsHlsMediaCustomTagApi(configs...)
    api.CustomTag = customTagApi
    typeApi, err := NewEncodingManifestsHlsMediaTypeApi(configs...)
    api.Type = typeApi
    videoApi, err := NewEncodingManifestsHlsMediaVideoApi(configs...)
    api.Video = videoApi
    audioApi, err := NewEncodingManifestsHlsMediaAudioApi(configs...)
    api.Audio = audioApi
    subtitlesApi, err := NewEncodingManifestsHlsMediaSubtitlesApi(configs...)
    api.Subtitles = subtitlesApi
    vttApi, err := NewEncodingManifestsHlsMediaVttApi(configs...)
    api.Vtt = vttApi
    closedCaptionsApi, err := NewEncodingManifestsHlsMediaClosedCaptionsApi(configs...)
    api.ClosedCaptions = closedCaptionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

