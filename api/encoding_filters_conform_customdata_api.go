package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersConformCustomdataAPI communicates with '/encoding/filters/conform/{filter_id}/customData' endpoints
type EncodingFiltersConformCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersConformCustomdataAPI constructor for EncodingFiltersConformCustomdataAPI that takes options as argument
func NewEncodingFiltersConformCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersConformCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersConformCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersConformCustomdataAPIWithClient constructor for EncodingFiltersConformCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersConformCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersConformCustomdataAPI {
	a := &EncodingFiltersConformCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Conform Filter Custom Data
func (api *EncodingFiltersConformCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/conform/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
