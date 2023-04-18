package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// StreamsLiveStartAPI communicates with '/streams/live/{stream_id}/start' endpoints
type StreamsLiveStartAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsLiveStartAPI constructor for StreamsLiveStartAPI that takes options as argument
func NewStreamsLiveStartAPI(options ...apiclient.APIClientOption) (*StreamsLiveStartAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsLiveStartAPIWithClient(apiClient), nil
}

// NewStreamsLiveStartAPIWithClient constructor for StreamsLiveStartAPI that takes an APIClient as argument
func NewStreamsLiveStartAPIWithClient(apiClient *apiclient.APIClient) *StreamsLiveStartAPI {
	a := &StreamsLiveStartAPI{apiClient: apiClient}
	return a
}

// Update Start live stream by id
func (api *StreamsLiveStartAPI) Update(streamId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	err := api.apiClient.Put("/streams/live/{stream_id}/start", nil, nil, reqParams)
	return err
}
