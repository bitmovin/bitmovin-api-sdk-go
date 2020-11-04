package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI constructor for EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI {
	a := &EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI{apiClient: apiClient}
	return a
}

// Get ClearKey DRM Custom Data of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmClearkeyCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/clearkey/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
