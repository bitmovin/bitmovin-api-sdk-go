package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI intermediary API object with no endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI struct {
	apiClient *apiclient.APIClient

	// Vtt communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/vtt' endpoints
	Vtt *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPI
	// Sprite communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/sprite' endpoints
	Sprite *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPI
	// Fmp4 communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4' endpoints
	Fmp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4API
	// ChunkedText communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/chunked-text' endpoints
	ChunkedText *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPI
	// Cmaf communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf' endpoints
	Cmaf *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI
	// Mp4 communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4' endpoints
	Mp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4API
	// Webm communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm' endpoints
	Webm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPI
	// ProgressiveWebm communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/progressive-webm' endpoints
	ProgressiveWebm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPI
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsAPI{apiClient: apiClient}
	a.Vtt = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttAPIWithClient(apiClient)
	a.Sprite = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSpriteAPIWithClient(apiClient)
	a.Fmp4 = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4APIWithClient(apiClient)
	a.ChunkedText = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsChunkedTextAPIWithClient(apiClient)
	a.Cmaf = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIWithClient(apiClient)
	a.Mp4 = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4APIWithClient(apiClient)
	a.Webm = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmAPIWithClient(apiClient)
	a.ProgressiveWebm = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsProgressiveWebmAPIWithClient(apiClient)

	return a
}
