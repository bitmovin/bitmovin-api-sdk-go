package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveAPI communicates with '/encoding/encodings/{encoding_id}/live' endpoints
type EncodingEncodingsLiveAPI struct {
	apiClient *apiclient.APIClient

	// ResetLiveManifestTimeshift communicates with '/encoding/encodings/{encoding_id}/live/reset-live-manifest-timeshift' endpoints
	ResetLiveManifestTimeshift *EncodingEncodingsLiveResetLiveManifestTimeshiftAPI
	// Heartbeat communicates with '/encoding/encodings/{encoding_id}/live/heartbeat' endpoints
	Heartbeat *EncodingEncodingsLiveHeartbeatAPI
	// Hd communicates with '/encoding/encodings/{encoding_id}/live/hd/start' endpoints
	Hd *EncodingEncodingsLiveHdAPI
	// InsertableContent communicates with '/encoding/encodings/{encoding_id}/live/insertable-content' endpoints
	InsertableContent *EncodingEncodingsLiveInsertableContentAPI
	// Scte35Cue communicates with '/encoding/encodings/{encoding_id}/live/scte-35-cue' endpoints
	Scte35Cue *EncodingEncodingsLiveScte35CueAPI
}

// NewEncodingEncodingsLiveAPI constructor for EncodingEncodingsLiveAPI that takes options as argument
func NewEncodingEncodingsLiveAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveAPIWithClient constructor for EncodingEncodingsLiveAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveAPI {
	a := &EncodingEncodingsLiveAPI{apiClient: apiClient}
	a.ResetLiveManifestTimeshift = NewEncodingEncodingsLiveResetLiveManifestTimeshiftAPIWithClient(apiClient)
	a.Heartbeat = NewEncodingEncodingsLiveHeartbeatAPIWithClient(apiClient)
	a.Hd = NewEncodingEncodingsLiveHdAPIWithClient(apiClient)
	a.InsertableContent = NewEncodingEncodingsLiveInsertableContentAPIWithClient(apiClient)
	a.Scte35Cue = NewEncodingEncodingsLiveScte35CueAPIWithClient(apiClient)

	return a
}

// Get Live Encoding Details
func (api *EncodingEncodingsLiveAPI) Get(encodingId string) (*model.LiveEncoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveEncoding
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetStartRequest Live Encoding Start Details
func (api *EncodingEncodingsLiveAPI) GetStartRequest(encodingId string) (*model.StartLiveEncodingRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.StartLiveEncodingRequest
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Restart Re-Start Live Encoding
func (api *EncodingEncodingsLiveAPI) Restart(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/restart", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start Live Encoding
func (api *EncodingEncodingsLiveAPI) Start(encodingId string, startLiveEncodingRequest model.StartLiveEncodingRequest) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/start", &startLiveEncodingRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Stop Live Encoding
func (api *EncodingEncodingsLiveAPI) Stop(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}
