package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersDeinterlaceCustomdataAPI communicates with '/encoding/filters/deinterlace/{filter_id}/customData' endpoints
type EncodingFiltersDeinterlaceCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersDeinterlaceCustomdataAPI constructor for EncodingFiltersDeinterlaceCustomdataAPI that takes options as argument
func NewEncodingFiltersDeinterlaceCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersDeinterlaceCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersDeinterlaceCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersDeinterlaceCustomdataAPIWithClient constructor for EncodingFiltersDeinterlaceCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersDeinterlaceCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersDeinterlaceCustomdataAPI {
	a := &EncodingFiltersDeinterlaceCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Deinterlace Filter Custom Data
func (api *EncodingFiltersDeinterlaceCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/filters/deinterlace/{filter_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
