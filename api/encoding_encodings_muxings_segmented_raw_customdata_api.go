package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsMuxingsSegmentedRawCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}/customData' endpoints
type EncodingEncodingsMuxingsSegmentedRawCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsMuxingsSegmentedRawCustomdataAPI constructor for EncodingEncodingsMuxingsSegmentedRawCustomdataAPI that takes options as argument
func NewEncodingEncodingsMuxingsSegmentedRawCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsSegmentedRawCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsSegmentedRawCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsSegmentedRawCustomdataAPIWithClient constructor for EncodingEncodingsMuxingsSegmentedRawCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsSegmentedRawCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsSegmentedRawCustomdataAPI {
	a := &EncodingEncodingsMuxingsSegmentedRawCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Segmented RAW muxing Custom Data
func (api *EncodingEncodingsMuxingsSegmentedRawCustomdataAPI) Get(encodingId string, muxingId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
