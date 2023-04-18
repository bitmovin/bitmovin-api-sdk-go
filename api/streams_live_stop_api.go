package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// StreamsLiveStopAPI communicates with '/streams/live/{stream_id}/stop' endpoints
type StreamsLiveStopAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsLiveStopAPI constructor for StreamsLiveStopAPI that takes options as argument
func NewStreamsLiveStopAPI(options ...apiclient.APIClientOption) (*StreamsLiveStopAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsLiveStopAPIWithClient(apiClient), nil
}

// NewStreamsLiveStopAPIWithClient constructor for StreamsLiveStopAPI that takes an APIClient as argument
func NewStreamsLiveStopAPIWithClient(apiClient *apiclient.APIClient) *StreamsLiveStopAPI {
	a := &StreamsLiveStopAPI{apiClient: apiClient}
	return a
}

// Update Stop live stream by id
func (api *StreamsLiveStopAPI) Update(streamId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	err := api.apiClient.Put("/streams/live/{stream_id}/stop", nil, nil, reqParams)
	return err
}
