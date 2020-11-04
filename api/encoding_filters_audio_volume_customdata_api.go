package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersAudioVolumeCustomdataAPI communicates with '/encoding/filters/audio-volume/{filter_id}/customData' endpoints
type EncodingFiltersAudioVolumeCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersAudioVolumeCustomdataAPI constructor for EncodingFiltersAudioVolumeCustomdataAPI that takes options as argument
func NewEncodingFiltersAudioVolumeCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAudioVolumeCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersAudioVolumeCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAudioVolumeCustomdataAPIWithClient constructor for EncodingFiltersAudioVolumeCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersAudioVolumeCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAudioVolumeCustomdataAPI {
	a := &EncodingFiltersAudioVolumeCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Audio Volume Filter Custom Data
func (api *EncodingFiltersAudioVolumeCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/audio-volume/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
