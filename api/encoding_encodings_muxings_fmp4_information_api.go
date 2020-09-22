package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsFmp4InformationAPI communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/information' endpoints
type EncodingEncodingsMuxingsFmp4InformationAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsFmp4InformationAPI constructor for EncodingEncodingsMuxingsFmp4InformationAPI that takes options as argument
func NewEncodingEncodingsMuxingsFmp4InformationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsFmp4InformationAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsFmp4InformationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsFmp4InformationAPIWithClient constructor for EncodingEncodingsMuxingsFmp4InformationAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsFmp4InformationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsFmp4InformationAPI {
	a := &EncodingEncodingsMuxingsFmp4InformationAPI{apiClient: apiClient}
	return a
}

// Get Fmp4 muxing Information
func (api *EncodingEncodingsMuxingsFmp4InformationAPI) Get(encodingId string, muxingId string) (*model.Fmp4MuxingInformation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.Fmp4MuxingInformation
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/information", nil, &responseModel, reqParams)
	return &responseModel, err
}
