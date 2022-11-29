package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// StreamsLiveAPI communicates with '/streams/live/{stream_id}' endpoints
type StreamsLiveAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsLiveAPI constructor for StreamsLiveAPI that takes options as argument
func NewStreamsLiveAPI(options ...apiclient.APIClientOption) (*StreamsLiveAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsLiveAPIWithClient(apiClient), nil
}

// NewStreamsLiveAPIWithClient constructor for StreamsLiveAPI that takes an APIClient as argument
func NewStreamsLiveAPIWithClient(apiClient *apiclient.APIClient) *StreamsLiveAPI {
	a := &StreamsLiveAPI{apiClient: apiClient}
	return a
}

// PatchStreamsLive Update stream by id
func (api *StreamsLiveAPI) PatchStreamsLive(streamId string, streamsLiveUpdateRequest model.StreamsLiveUpdateRequest) (*model.StreamsLiveUpdateRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamsLiveUpdateRequest
	err := api.apiClient.Patch("/streams/live/{stream_id}", &streamsLiveUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}
