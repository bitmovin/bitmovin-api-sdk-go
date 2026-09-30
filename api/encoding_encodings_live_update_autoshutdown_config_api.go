package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveUpdateAutoshutdownConfigAPI communicates with '/encoding/encodings/{encoding_id}/live/update-autoshutdown-config' endpoints
type EncodingEncodingsLiveUpdateAutoshutdownConfigAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveUpdateAutoshutdownConfigAPI constructor for EncodingEncodingsLiveUpdateAutoshutdownConfigAPI that takes options as argument
func NewEncodingEncodingsLiveUpdateAutoshutdownConfigAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveUpdateAutoshutdownConfigAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveUpdateAutoshutdownConfigAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveUpdateAutoshutdownConfigAPIWithClient constructor for EncodingEncodingsLiveUpdateAutoshutdownConfigAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveUpdateAutoshutdownConfigAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveUpdateAutoshutdownConfigAPI {
	a := &EncodingEncodingsLiveUpdateAutoshutdownConfigAPI{apiClient: apiClient}
	return a
}

// Create Replace Live Auto Shutdown Configuration
func (api *EncodingEncodingsLiveUpdateAutoshutdownConfigAPI) Create(encodingId string, liveAutoShutdownConfigurationUpdateRequest model.LiveAutoShutdownConfigurationUpdateRequest) (*model.LiveAutoShutdownConfigurationUpdateResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveAutoShutdownConfigurationUpdateResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/update-autoshutdown-config", &liveAutoShutdownConfigurationUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}
