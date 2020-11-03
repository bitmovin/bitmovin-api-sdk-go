package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingManifestsDashPeriodsAdaptationsetsAPI intermediary API object with no endpoints
type EncodingManifestsDashPeriodsAdaptationsetsAPI struct {
    apiClient *apiclient.APIClient

    // Audio communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/audio' endpoints
    Audio *EncodingManifestsDashPeriodsAdaptationsetsAudioAPI
    // Video communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video' endpoints
    Video *EncodingManifestsDashPeriodsAdaptationsetsVideoAPI
    // Subtitle communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle' endpoints
    Subtitle *EncodingManifestsDashPeriodsAdaptationsetsSubtitleAPI
    // Representations intermediary API object with no endpoints
    Representations *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI
    // Contentprotection communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/contentprotection' endpoints
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPI

}

// NewEncodingManifestsDashPeriodsAdaptationsetsAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsAPI {
    a := &EncodingManifestsDashPeriodsAdaptationsetsAPI{apiClient: apiClient}
    a.Audio = NewEncodingManifestsDashPeriodsAdaptationsetsAudioAPIWithClient(apiClient)
    a.Video = NewEncodingManifestsDashPeriodsAdaptationsetsVideoAPIWithClient(apiClient)
    a.Subtitle = NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleAPIWithClient(apiClient)
    a.Representations = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient(apiClient)
    a.Contentprotection = NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionAPIWithClient(apiClient)

    return a
}


