package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsInputStreamsSubtitlesAPI intermediary API object with no endpoints
type EncodingEncodingsInputStreamsSubtitlesAPI struct {
    apiClient *apiclient.APIClient

    // DvbSubtitle communicates with '/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle' endpoints
    DvbSubtitle *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI
    // DvbTeletext communicates with '/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext' endpoints
    DvbTeletext *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI

}

// NewEncodingEncodingsInputStreamsSubtitlesAPI constructor for EncodingEncodingsInputStreamsSubtitlesAPI that takes options as argument
func NewEncodingEncodingsInputStreamsSubtitlesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsSubtitlesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsSubtitlesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsSubtitlesAPIWithClient constructor for EncodingEncodingsInputStreamsSubtitlesAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsSubtitlesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsSubtitlesAPI {
    a := &EncodingEncodingsInputStreamsSubtitlesAPI{apiClient: apiClient}
    a.DvbSubtitle = NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIWithClient(apiClient)
    a.DvbTeletext = NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIWithClient(apiClient)

    return a
}


