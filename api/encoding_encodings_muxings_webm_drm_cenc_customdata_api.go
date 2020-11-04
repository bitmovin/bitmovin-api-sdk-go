package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsWebmDrmCencCustomdataAPI constructor for EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmDrmCencCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmDrmCencCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmDrmCencCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmDrmCencCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI {
	a := &EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI{apiClient: apiClient}
	return a
}

// Get CENC DRM Custom Data of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmCencCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
