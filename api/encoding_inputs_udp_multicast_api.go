package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsUdpMulticastAPI communicates with '/encoding/inputs/udp-multicast' endpoints
type EncodingInputsUdpMulticastAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/udp-multicast/{input_id}/customData' endpoints
	Customdata *EncodingInputsUdpMulticastCustomdataAPI
}

// NewEncodingInputsUdpMulticastAPI constructor for EncodingInputsUdpMulticastAPI that takes options as argument
func NewEncodingInputsUdpMulticastAPI(options ...apiclient.APIClientOption) (*EncodingInputsUdpMulticastAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsUdpMulticastAPIWithClient(apiClient), nil
}

// NewEncodingInputsUdpMulticastAPIWithClient constructor for EncodingInputsUdpMulticastAPI that takes an APIClient as argument
func NewEncodingInputsUdpMulticastAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsUdpMulticastAPI {
	a := &EncodingInputsUdpMulticastAPI{apiClient: apiClient}
	a.Customdata = NewEncodingInputsUdpMulticastCustomdataAPIWithClient(apiClient)

	return a
}

// Create UDP multicast input
func (api *EncodingInputsUdpMulticastAPI) Create(udpMulticastInput model.UdpMulticastInput) (*model.UdpMulticastInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.UdpMulticastInput
	err := api.apiClient.Post("/encoding/inputs/udp-multicast", &udpMulticastInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete UDP multicast input
func (api *EncodingInputsUdpMulticastAPI) Delete(inputId string) (*model.UdpMulticastInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.UdpMulticastInput
	err := api.apiClient.Delete("/encoding/inputs/udp-multicast/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get UDP multicast Input Details
func (api *EncodingInputsUdpMulticastAPI) Get(inputId string) (*model.UdpMulticastInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.UdpMulticastInput
	err := api.apiClient.Get("/encoding/inputs/udp-multicast/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List UDP multicast inputs
func (api *EncodingInputsUdpMulticastAPI) List(queryParams ...func(*EncodingInputsUdpMulticastAPIListQueryParams)) (*pagination.UdpMulticastInputsListPagination, error) {
	queryParameters := &EncodingInputsUdpMulticastAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.UdpMulticastInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/udp-multicast", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsUdpMulticastAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsUdpMulticastAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsUdpMulticastAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
