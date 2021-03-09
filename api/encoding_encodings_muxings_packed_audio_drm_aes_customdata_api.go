package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes/{drm_id}/customData' endpoints
type EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI constructor for EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI {
	a := &EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI{apiClient: apiClient}
	return a
}

// Get AES encryption Custom Data of a Packed Audio muxing
func (api *EncodingEncodingsMuxingsPackedAudioDrmAesCustomdataAPI) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
		params.PathParams["drm_id"] = drmId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/drm/aes/{drm_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
