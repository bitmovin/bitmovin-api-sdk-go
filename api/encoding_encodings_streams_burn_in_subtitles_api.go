package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsStreamsBurnInSubtitlesAPI intermediary API object with no endpoints
type EncodingEncodingsStreamsBurnInSubtitlesAPI struct {
    apiClient *apiclient.APIClient

    // Dvbsub communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub' endpoints
    Dvbsub *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI
    // Srt communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt' endpoints
    Srt *EncodingEncodingsStreamsBurnInSubtitlesSrtAPI

}

// NewEncodingEncodingsStreamsBurnInSubtitlesAPI constructor for EncodingEncodingsStreamsBurnInSubtitlesAPI that takes options as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsBurnInSubtitlesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsBurnInSubtitlesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsBurnInSubtitlesAPIWithClient constructor for EncodingEncodingsStreamsBurnInSubtitlesAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsBurnInSubtitlesAPI {
    a := &EncodingEncodingsStreamsBurnInSubtitlesAPI{apiClient: apiClient}
    a.Dvbsub = NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIWithClient(apiClient)
    a.Srt = NewEncodingEncodingsStreamsBurnInSubtitlesSrtAPIWithClient(apiClient)

    return a
}


