package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveWavInformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsProgressiveWavInformationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveWavInformationAPI constructor for EncodingEncodingsMuxingsProgressiveWavInformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWavInformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWavInformationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWavInformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWavInformationAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWavInformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWavInformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWavInformationAPI {
	a := &EncodingEncodingsMuxingsProgressiveWavInformationAPI{apiClient: apiClient}
	return a
}

// Get Progressive WAV muxing Information
func (api *EncodingEncodingsMuxingsProgressiveWavInformationAPI) Get(encodingId string, muxingId string) (*model.ProgressiveWavMuxingInformation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ProgressiveWavMuxingInformation
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-wav/{muxing_id}/information", nil, &responseModel, reqParams)
	return &responseModel, err
}
