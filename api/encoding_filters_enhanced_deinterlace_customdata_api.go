package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersEnhancedDeinterlaceCustomdataAPI communicates with '/encoding/filters/enhanced-deinterlace/{filter_id}/customData' endpoints
type EncodingFiltersEnhancedDeinterlaceCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersEnhancedDeinterlaceCustomdataAPI constructor for EncodingFiltersEnhancedDeinterlaceCustomdataAPI that takes options as argument
func NewEncodingFiltersEnhancedDeinterlaceCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersEnhancedDeinterlaceCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersEnhancedDeinterlaceCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersEnhancedDeinterlaceCustomdataAPIWithClient constructor for EncodingFiltersEnhancedDeinterlaceCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersEnhancedDeinterlaceCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersEnhancedDeinterlaceCustomdataAPI {
	a := &EncodingFiltersEnhancedDeinterlaceCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Enhanced Deinterlace Filter Custom Data
func (api *EncodingFiltersEnhancedDeinterlaceCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/enhanced-deinterlace/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
