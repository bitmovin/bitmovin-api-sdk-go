package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersRotateCustomdataAPI communicates with '/encoding/filters/rotate/{filter_id}/customData' endpoints
type EncodingFiltersRotateCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersRotateCustomdataAPI constructor for EncodingFiltersRotateCustomdataAPI that takes options as argument
func NewEncodingFiltersRotateCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersRotateCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersRotateCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersRotateCustomdataAPIWithClient constructor for EncodingFiltersRotateCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersRotateCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersRotateCustomdataAPI {
	a := &EncodingFiltersRotateCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Rotate Filter Custom Data
func (api *EncodingFiltersRotateCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/rotate/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
