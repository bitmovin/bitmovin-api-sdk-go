package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi struct {
    apiClient *common.ApiClient
    Vtt *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttApi
    Sidecar *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsSidecarApi
    Fmp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4Api
    Mp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4Api
    Webm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi{apiClient: apiClient}

    vttApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttApi(configs...)
    api.Vtt = vttApi
    sidecarApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsSidecarApi(configs...)
    api.Sidecar = sidecarApi
    fmp4Api, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4Api(configs...)
    api.Fmp4 = fmp4Api
    mp4Api, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4Api(configs...)
    api.Mp4 = mp4Api
    webmApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi(configs...)
    api.Webm = webmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

