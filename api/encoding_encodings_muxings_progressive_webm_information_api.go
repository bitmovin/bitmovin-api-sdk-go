package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveWebmInformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsProgressiveWebmInformationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveWebmInformationAPI constructor for EncodingEncodingsMuxingsProgressiveWebmInformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveWebmInformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveWebmInformationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveWebmInformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveWebmInformationAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveWebmInformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveWebmInformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveWebmInformationAPI {
	a := &EncodingEncodingsMuxingsProgressiveWebmInformationAPI{apiClient: apiClient}
	return a
}

// Get Progressive WebM muxing Information
func (api *EncodingEncodingsMuxingsProgressiveWebmInformationAPI) Get(encodingId string, muxingId string) (*model.ProgressiveWebmMuxingInformation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ProgressiveWebmMuxingInformation
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/information", nil, &responseModel, reqParams)
	return &responseModel, err
}
