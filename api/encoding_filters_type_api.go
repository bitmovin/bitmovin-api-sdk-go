package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersTypeAPI communicates with '/encoding/filters/{filter_id}/type' endpoints
type EncodingFiltersTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingFiltersTypeAPI constructor for EncodingFiltersTypeAPI that takes options as argument
func NewEncodingFiltersTypeAPI(options ...apiclient.APIClientOption) (*EncodingFiltersTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersTypeAPIWithClient(apiClient), nil
}

// NewEncodingFiltersTypeAPIWithClient constructor for EncodingFiltersTypeAPI that takes an APIClient as argument
func NewEncodingFiltersTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersTypeAPI {
	a := &EncodingFiltersTypeAPI{apiClient: apiClient}
	return a
}

// Get Filter Type
func (api *EncodingFiltersTypeAPI) Get(filterId string) (*model.FilterType, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.FilterType
	err := api.apiClient.Get("/encoding/filters/{filter_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
