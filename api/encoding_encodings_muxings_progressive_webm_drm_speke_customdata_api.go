package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI{apiClient: apiClient}
	return a
}

// Get SPEKE DRM Custom Data of a Progressive WebM muxing
func (api *EncodingEncodingsMuxingsProgressiveWebmDrmSpekeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
