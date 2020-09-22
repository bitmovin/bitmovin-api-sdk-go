package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI{apiClient: apiClient}
	return a
}

// Get PlayReady DRM Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
