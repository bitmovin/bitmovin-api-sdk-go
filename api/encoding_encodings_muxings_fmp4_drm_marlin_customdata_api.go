package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Marlin DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmMarlinCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/marlin/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
