package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsProgressiveTsInformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsProgressiveTsInformationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsProgressiveTsInformationAPI constructor for EncodingEncodingsMuxingsProgressiveTsInformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsInformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsInformationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsProgressiveTsInformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsInformationAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsInformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsInformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsInformationAPI {
	a := &EncodingEncodingsMuxingsProgressiveTsInformationAPI{apiClient: apiClient}
	return a
}

// Get Progressive TS muxing Information
func (api *EncodingEncodingsMuxingsProgressiveTsInformationAPI) Get(encodingId string, muxingId string) (*model.ProgressiveTsMuxingInformation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ProgressiveTsMuxingInformation
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/information", nil, &responseModel, reqParams)
	return &responseModel, err
}
