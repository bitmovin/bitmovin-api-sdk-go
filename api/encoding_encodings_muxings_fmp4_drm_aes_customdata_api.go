package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI constructor for EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4DrmAesCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI {
	a := &EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI{apiClient: apiClient}
	return a
}

// Get AES encryption Custom Data of an fMP4 muxing
func (api *EncodingEncodingsMuxingsFmp4DrmAesCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/aes/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
