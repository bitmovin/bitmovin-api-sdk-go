package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker/{nexguard_id}/customData' endpoints
type EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI constructor for EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPIWithClient constructor for EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI {
	a := &EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Nexguard file marker watermarking configuration Custom Data
func (api *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI) Get(encodingId string, streamId string, nexguardId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["nexguard_id"] = nexguardId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker/{nexguard_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
