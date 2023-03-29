package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// StreamsConfigsAPI communicates with '/streams/configs/{config_id}' endpoints
type StreamsConfigsAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsConfigsAPI constructor for StreamsConfigsAPI that takes options as argument
func NewStreamsConfigsAPI(options ...apiclient.APIClientOption) (*StreamsConfigsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsConfigsAPIWithClient(apiClient), nil
}

// NewStreamsConfigsAPIWithClient constructor for StreamsConfigsAPI that takes an APIClient as argument
func NewStreamsConfigsAPIWithClient(apiClient *apiclient.APIClient) *StreamsConfigsAPI {
	a := &StreamsConfigsAPI{apiClient: apiClient}
	return a
}

// PatchStreamConfig Update stream config by id
func (api *StreamsConfigsAPI) PatchStreamConfig(configId string, streamsConfigUpdateRequest model.StreamsConfigUpdateRequest) (*model.StreamsConfigResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["config_id"] = configId
	}

	var responseModel model.StreamsConfigResponse
	err := api.apiClient.Patch("/streams/configs/{config_id}", &streamsConfigUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}
