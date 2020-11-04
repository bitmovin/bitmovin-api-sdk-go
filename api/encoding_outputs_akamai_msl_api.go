package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsAkamaiMslAPI communicates with '/encoding/outputs/akamai-msl' endpoints
type EncodingOutputsAkamaiMslAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/akamai-msl/{output_id}/customData' endpoints
	Customdata *EncodingOutputsAkamaiMslCustomdataAPI
}

// NewEncodingOutputsAkamaiMslAPI constructor for EncodingOutputsAkamaiMslAPI that takes options as argument
func NewEncodingOutputsAkamaiMslAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAkamaiMslAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsAkamaiMslAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAkamaiMslAPIWithClient constructor for EncodingOutputsAkamaiMslAPI that takes an APIClient as argument
func NewEncodingOutputsAkamaiMslAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAkamaiMslAPI {
	a := &EncodingOutputsAkamaiMslAPI{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsAkamaiMslCustomdataAPIWithClient(apiClient)

	return a
}

// Create Akamai MSL Output
func (api *EncodingOutputsAkamaiMslAPI) Create(akamaiMslOutput model.AkamaiMslOutput) (*model.AkamaiMslOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AkamaiMslOutput
	err := api.apiClient.Post("/encoding/outputs/akamai-msl", &akamaiMslOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Akamai MSL Output
func (api *EncodingOutputsAkamaiMslAPI) Delete(outputId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/outputs/akamai-msl/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Akamai MSL Output Details
func (api *EncodingOutputsAkamaiMslAPI) Get(outputId string) (*model.AkamaiMslOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.AkamaiMslOutput
	err := api.apiClient.Get("/encoding/outputs/akamai-msl/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Akamai MSL Outputs
func (api *EncodingOutputsAkamaiMslAPI) List(queryParams ...func(*EncodingOutputsAkamaiMslAPIListQueryParams)) (*pagination.AkamaiMslOutputsListPagination, error) {
	queryParameters := &EncodingOutputsAkamaiMslAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AkamaiMslOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/akamai-msl", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsAkamaiMslAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsAkamaiMslAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsAkamaiMslAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
