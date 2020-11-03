package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsInputStreamsTrimmingAPI intermediary API object with no endpoints
type EncodingEncodingsInputStreamsTrimmingAPI struct {
    apiClient *apiclient.APIClient

    // TimeBased communicates with '/encoding/encodings/{encoding_id}/input-streams/trimming/time-based' endpoints
    TimeBased *EncodingEncodingsInputStreamsTrimmingTimeBasedAPI
    // TimecodeTrack communicates with '/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track' endpoints
    TimecodeTrack *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI
    // H264PictureTiming communicates with '/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing' endpoints
    H264PictureTiming *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI

}

// NewEncodingEncodingsInputStreamsTrimmingAPI constructor for EncodingEncodingsInputStreamsTrimmingAPI that takes options as argument
func NewEncodingEncodingsInputStreamsTrimmingAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsTrimmingAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsTrimmingAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsTrimmingAPIWithClient constructor for EncodingEncodingsInputStreamsTrimmingAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsTrimmingAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsTrimmingAPI {
    a := &EncodingEncodingsInputStreamsTrimmingAPI{apiClient: apiClient}
    a.TimeBased = NewEncodingEncodingsInputStreamsTrimmingTimeBasedAPIWithClient(apiClient)
    a.TimecodeTrack = NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIWithClient(apiClient)
    a.H264PictureTiming = NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIWithClient(apiClient)

    return a
}


