package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI{apiClient: apiClient}
	return a
}

// Get FairPlay DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
