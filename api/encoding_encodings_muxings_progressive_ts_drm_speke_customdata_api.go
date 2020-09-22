package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI constructor for EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI{apiClient: apiClient}
	return a
}

// Get SPEKE DRM Custom Data of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
