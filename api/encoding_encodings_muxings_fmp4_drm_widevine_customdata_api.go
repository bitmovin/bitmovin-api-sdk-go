package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Widevine DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmWidevineCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/widevine/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
