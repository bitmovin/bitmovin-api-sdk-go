package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingLiveEncodingsActionsAPI communicates with '/encoding/live/encodings/{encoding_id}/actions/update-rtmp-ingest-points' endpoints
type EncodingLiveEncodingsActionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveEncodingsActionsAPI constructor for EncodingLiveEncodingsActionsAPI that takes options as argument
func NewEncodingLiveEncodingsActionsAPI(options ...apiclient.APIClientOption) (*EncodingLiveEncodingsActionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveEncodingsActionsAPIWithClient(apiClient), nil
}

// NewEncodingLiveEncodingsActionsAPIWithClient constructor for EncodingLiveEncodingsActionsAPI that takes an APIClient as argument
func NewEncodingLiveEncodingsActionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveEncodingsActionsAPI {
	a := &EncodingLiveEncodingsActionsAPI{apiClient: apiClient}
	return a
}

// Patch Update the ingest points of a Redundant RTMP Input
func (api *EncodingLiveEncodingsActionsAPI) Patch(encodingId string, updateEncodingRtmpIngestPointRequest model.UpdateEncodingRtmpIngestPointRequest) (*model.UpdateEncodingRtmpIngestPointResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.UpdateEncodingRtmpIngestPointResponse
	err := api.apiClient.Patch("/encoding/live/encodings/{encoding_id}/actions/update-rtmp-ingest-points", &updateEncodingRtmpIngestPointRequest, &responseModel, reqParams)
	return &responseModel, err
}
