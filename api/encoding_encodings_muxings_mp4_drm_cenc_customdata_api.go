package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp4DrmCencCustomdataAPI constructor for EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp4DrmCencCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp4DrmCencCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp4DrmCencCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp4DrmCencCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI {
	a := &EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI{apiClient: apiClient}
	return a
}

// Get CENC DRM Custom Data of an MP4 muxing
func (api *EncodingEncodingsMuxingsMp4DrmCencCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/cenc/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
