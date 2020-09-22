package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsInputStreamsTypeAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}/type' endpoints
type EncodingEncodingsInputStreamsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsTypeAPI constructor for EncodingEncodingsInputStreamsTypeAPI that takes options as argument
func NewEncodingEncodingsInputStreamsTypeAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsTypeAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsTypeAPIWithClient constructor for EncodingEncodingsInputStreamsTypeAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsTypeAPI {
	a := &EncodingEncodingsInputStreamsTypeAPI{apiClient: apiClient}
	return a
}

// Get Input Stream Type
func (api *EncodingEncodingsInputStreamsTypeAPI) Get(encodingId string, inputStreamId string) (*model.InputStreamTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.InputStreamTypeResponse
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
