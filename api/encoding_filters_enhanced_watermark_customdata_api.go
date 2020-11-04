package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersEnhancedWatermarkCustomdataAPI communicates with '/encoding/filters/enhanced-watermark/{filter_id}/customData' endpoints
type EncodingFiltersEnhancedWatermarkCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersEnhancedWatermarkCustomdataAPI constructor for EncodingFiltersEnhancedWatermarkCustomdataAPI that takes options as argument
func NewEncodingFiltersEnhancedWatermarkCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersEnhancedWatermarkCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersEnhancedWatermarkCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersEnhancedWatermarkCustomdataAPIWithClient constructor for EncodingFiltersEnhancedWatermarkCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersEnhancedWatermarkCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersEnhancedWatermarkCustomdataAPI {
	a := &EncodingFiltersEnhancedWatermarkCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Enhanced Watermark Filter Custom Data
func (api *EncodingFiltersEnhancedWatermarkCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/enhanced-watermark/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
