package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersWatermarkCustomdataAPI communicates with '/encoding/filters/watermark/{filter_id}/customData' endpoints
type EncodingFiltersWatermarkCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersWatermarkCustomdataAPI constructor for EncodingFiltersWatermarkCustomdataAPI that takes options as argument
func NewEncodingFiltersWatermarkCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersWatermarkCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersWatermarkCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersWatermarkCustomdataAPIWithClient constructor for EncodingFiltersWatermarkCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersWatermarkCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersWatermarkCustomdataAPI {
	a := &EncodingFiltersWatermarkCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Watermark Filter Custom Data
func (api *EncodingFiltersWatermarkCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/watermark/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
