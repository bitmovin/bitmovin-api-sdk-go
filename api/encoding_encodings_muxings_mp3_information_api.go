package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsMp3InformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsMp3InformationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsMp3InformationAPI constructor for EncodingEncodingsMuxingsMp3InformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsMp3InformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsMp3InformationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsMp3InformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsMp3InformationAPIWithClient constructor for EncodingEncodingsMuxingsMp3InformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsMp3InformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsMp3InformationAPI {
	a := &EncodingEncodingsMuxingsMp3InformationAPI{apiClient: apiClient}
	return a
}

// Get MP3 muxing Information
func (api *EncodingEncodingsMuxingsMp3InformationAPI) Get(encodingId string, muxingId string) (*model.Mp3MuxingInformation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.Mp3MuxingInformation
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/information", nil, &responseModel, reqParams)
	return &responseModel, err
}
