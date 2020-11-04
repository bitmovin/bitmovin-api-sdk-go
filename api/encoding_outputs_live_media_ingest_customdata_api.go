package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsLiveMediaIngestCustomdataAPI communicates with '/encoding/outputs/live-media-ingest/{output_id}/customData' endpoints
type EncodingOutputsLiveMediaIngestCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsLiveMediaIngestCustomdataAPI constructor for EncodingOutputsLiveMediaIngestCustomdataAPI that takes options as argument
func NewEncodingOutputsLiveMediaIngestCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsLiveMediaIngestCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsLiveMediaIngestCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsLiveMediaIngestCustomdataAPIWithClient constructor for EncodingOutputsLiveMediaIngestCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsLiveMediaIngestCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsLiveMediaIngestCustomdataAPI {
	a := &EncodingOutputsLiveMediaIngestCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Live Media Ingest Output Custom Data
func (api *EncodingOutputsLiveMediaIngestCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/outputs/live-media-ingest/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
