package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsPackedAudioInformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsPackedAudioInformationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsPackedAudioInformationAPI constructor for EncodingEncodingsMuxingsPackedAudioInformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsPackedAudioInformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsPackedAudioInformationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsPackedAudioInformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsPackedAudioInformationAPIWithClient constructor for EncodingEncodingsMuxingsPackedAudioInformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsPackedAudioInformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsPackedAudioInformationAPI {
	a := &EncodingEncodingsMuxingsPackedAudioInformationAPI{apiClient: apiClient}
	return a
}

// Get Packed Audio muxing Information
func (api *EncodingEncodingsMuxingsPackedAudioInformationAPI) Get(encodingId string, muxingId string) (*model.PackedAudioMuxingInformation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.PackedAudioMuxingInformation
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/packed-audio/{muxing_id}/information", nil, &responseModel, reqParams)
	return &responseModel, err
}
