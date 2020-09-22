package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersCropCustomdataAPI communicates with '/encoding/filters/crop/{filter_id}/customData' endpoints
type EncodingFiltersCropCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersCropCustomdataAPI constructor for EncodingFiltersCropCustomdataAPI that takes options as argument
func NewEncodingFiltersCropCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersCropCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersCropCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersCropCustomdataAPIWithClient constructor for EncodingFiltersCropCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersCropCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersCropCustomdataAPI {
	a := &EncodingFiltersCropCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Crop Filter Custom Data
func (api *EncodingFiltersCropCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/crop/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
