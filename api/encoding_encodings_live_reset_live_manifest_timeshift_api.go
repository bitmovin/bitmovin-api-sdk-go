package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveResetLiveManifestTimeshiftAPI communicates with '/encoding/encodings/{encoding_id}/live/reset-live-manifest-timeshift' endpoints
type EncodingEncodingsLiveResetLiveManifestTimeshiftAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveResetLiveManifestTimeshiftAPI constructor for EncodingEncodingsLiveResetLiveManifestTimeshiftAPI that takes options as argument
func NewEncodingEncodingsLiveResetLiveManifestTimeshiftAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveResetLiveManifestTimeshiftAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveResetLiveManifestTimeshiftAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveResetLiveManifestTimeshiftAPIWithClient constructor for EncodingEncodingsLiveResetLiveManifestTimeshiftAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveResetLiveManifestTimeshiftAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveResetLiveManifestTimeshiftAPI {
	a := &EncodingEncodingsLiveResetLiveManifestTimeshiftAPI{apiClient: apiClient}
	return a
}

// Create Reset Live manifest time-shift
func (api *EncodingEncodingsLiveResetLiveManifestTimeshiftAPI) Create(encodingId string, resetLiveManifestTimeShift model.ResetLiveManifestTimeShift) (*model.ResetLiveManifestTimeShift, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ResetLiveManifestTimeShift
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/reset-live-manifest-timeshift", &resetLiveManifestTimeShift, &responseModel, reqParams)
	return &responseModel, err
}
