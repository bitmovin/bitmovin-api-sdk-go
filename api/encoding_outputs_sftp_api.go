package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsSftpAPI communicates with '/encoding/outputs/sftp' endpoints
type EncodingOutputsSftpAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/sftp/{output_id}/customData' endpoints
	Customdata *EncodingOutputsSftpCustomdataAPI
}

// NewEncodingOutputsSftpAPI constructor for EncodingOutputsSftpAPI that takes options as argument
func NewEncodingOutputsSftpAPI(options ...apiclient.APIClientOption) (*EncodingOutputsSftpAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsSftpAPIWithClient(apiClient), nil
}

// NewEncodingOutputsSftpAPIWithClient constructor for EncodingOutputsSftpAPI that takes an APIClient as argument
func NewEncodingOutputsSftpAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsSftpAPI {
	a := &EncodingOutputsSftpAPI{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsSftpCustomdataAPIWithClient(apiClient)

	return a
}

// Create SFTP Output
func (api *EncodingOutputsSftpAPI) Create(sftpOutput model.SftpOutput) (*model.SftpOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.SftpOutput
	err := api.apiClient.Post("/encoding/outputs/sftp", &sftpOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete SFTP Output
func (api *EncodingOutputsSftpAPI) Delete(outputId string) (*model.SftpOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.SftpOutput
	err := api.apiClient.Delete("/encoding/outputs/sftp/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get SFTP Output Details
func (api *EncodingOutputsSftpAPI) Get(outputId string) (*model.SftpOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.SftpOutput
	err := api.apiClient.Get("/encoding/outputs/sftp/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List SFTP Outputs
func (api *EncodingOutputsSftpAPI) List(queryParams ...func(*EncodingOutputsSftpAPIListQueryParams)) (*pagination.SftpOutputsListPagination, error) {
	queryParameters := &EncodingOutputsSftpAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SftpOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/sftp", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsSftpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsSftpAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsSftpAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
