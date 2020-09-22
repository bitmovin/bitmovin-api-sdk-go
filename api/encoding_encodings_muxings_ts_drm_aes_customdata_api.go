package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsTsDrmAesCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsTsDrmAesCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsTsDrmAesCustomdataAPI constructor for EncodingEncodingsMuxingsTsDrmAesCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsTsDrmAesCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsTsDrmAesCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsTsDrmAesCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsTsDrmAesCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsTsDrmAesCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsTsDrmAesCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsTsDrmAesCustomdataAPI {
	a := &EncodingEncodingsMuxingsTsDrmAesCustomdataAPI{apiClient: apiClient}
	return a
}

// Get AES encryption Custom Data of a TS muxing
func (api *EncodingEncodingsMuxingsTsDrmAesCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
