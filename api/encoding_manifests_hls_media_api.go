package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingManifestsHlsMediaAPI intermediary API object with no endpoints
type EncodingManifestsHlsMediaAPI struct {
    apiClient *apiclient.APIClient

    // CustomTags communicates with '/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags' endpoints
    CustomTags *EncodingManifestsHlsMediaCustomTagsAPI
    // Type communicates with '/encoding/manifests/hls/{manifest_id}/media/{media_id}/type' endpoints
    Type *EncodingManifestsHlsMediaTypeAPI
    // Video communicates with '/encoding/manifests/hls/{manifest_id}/media/video' endpoints
    Video *EncodingManifestsHlsMediaVideoAPI
    // Audio communicates with '/encoding/manifests/hls/{manifest_id}/media/audio' endpoints
    Audio *EncodingManifestsHlsMediaAudioAPI
    // Subtitles communicates with '/encoding/manifests/hls/{manifest_id}/media/subtitles' endpoints
    Subtitles *EncodingManifestsHlsMediaSubtitlesAPI
    // Vtt communicates with '/encoding/manifests/hls/{manifest_id}/media/vtt' endpoints
    Vtt *EncodingManifestsHlsMediaVttAPI
    // ClosedCaptions communicates with '/encoding/manifests/hls/{manifest_id}/media/closed-captions' endpoints
    ClosedCaptions *EncodingManifestsHlsMediaClosedCaptionsAPI

}

// NewEncodingManifestsHlsMediaAPI constructor for EncodingManifestsHlsMediaAPI that takes options as argument
func NewEncodingManifestsHlsMediaAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsHlsMediaAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaAPIWithClient constructor for EncodingManifestsHlsMediaAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaAPI {
    a := &EncodingManifestsHlsMediaAPI{apiClient: apiClient}
    a.CustomTags = NewEncodingManifestsHlsMediaCustomTagsAPIWithClient(apiClient)
    a.Type = NewEncodingManifestsHlsMediaTypeAPIWithClient(apiClient)
    a.Video = NewEncodingManifestsHlsMediaVideoAPIWithClient(apiClient)
    a.Audio = NewEncodingManifestsHlsMediaAudioAPIWithClient(apiClient)
    a.Subtitles = NewEncodingManifestsHlsMediaSubtitlesAPIWithClient(apiClient)
    a.Vtt = NewEncodingManifestsHlsMediaVttAPIWithClient(apiClient)
    a.ClosedCaptions = NewEncodingManifestsHlsMediaClosedCaptionsAPIWithClient(apiClient)

    return a
}


