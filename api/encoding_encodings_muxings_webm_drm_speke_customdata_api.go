package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI constructor for EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI {
	a := &EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI{apiClient: apiClient}
	return a
}

// Get SPEKE DRM Custom Data of a WebM muxing
func (api *EncodingEncodingsMuxingsWebmDrmSpekeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
