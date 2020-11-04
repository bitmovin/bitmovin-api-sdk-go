package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersEbuR128SinglePassCustomdataAPI communicates with '/encoding/filters/ebu-r128-single-pass/{filter_id}/customData' endpoints
type EncodingFiltersEbuR128SinglePassCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersEbuR128SinglePassCustomdataAPI constructor for EncodingFiltersEbuR128SinglePassCustomdataAPI that takes options as argument
func NewEncodingFiltersEbuR128SinglePassCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersEbuR128SinglePassCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersEbuR128SinglePassCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersEbuR128SinglePassCustomdataAPIWithClient constructor for EncodingFiltersEbuR128SinglePassCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersEbuR128SinglePassCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersEbuR128SinglePassCustomdataAPI {
	a := &EncodingFiltersEbuR128SinglePassCustomdataAPI{apiClient: apiClient}
	return a
}

// Get EBU R128 Single Pass Filter Custom Data
func (api *EncodingFiltersEbuR128SinglePassCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/ebu-r128-single-pass/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
