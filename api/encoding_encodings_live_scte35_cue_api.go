package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveScte35CueAPI communicates with '/encoding/encodings/{encoding_id}/live/scte-35-cue' endpoints
type EncodingEncodingsLiveScte35CueAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveScte35CueAPI constructor for EncodingEncodingsLiveScte35CueAPI that takes options as argument
func NewEncodingEncodingsLiveScte35CueAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveScte35CueAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveScte35CueAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveScte35CueAPIWithClient constructor for EncodingEncodingsLiveScte35CueAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveScte35CueAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveScte35CueAPI {
	a := &EncodingEncodingsLiveScte35CueAPI{apiClient: apiClient}
	return a
}

// Create Insert cue duration
func (api *EncodingEncodingsLiveScte35CueAPI) Create(encodingId string, scte35Cue model.Scte35Cue) (*model.Scte35Cue, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Scte35Cue
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/scte-35-cue", &scte35Cue, &responseModel, reqParams)
	return &responseModel, err
}
