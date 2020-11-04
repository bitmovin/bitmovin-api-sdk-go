package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI{apiClient: apiClient}
	return a
}

// Get SPEKE DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmSpekeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
