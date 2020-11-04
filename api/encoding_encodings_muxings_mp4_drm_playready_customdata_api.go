package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI constructor for EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI {
	a := &EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI{apiClient: apiClient}
	return a
}

// Get PlayReady DRM Custom Data of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
