package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsFtpAPI communicates with '/encoding/outputs/ftp' endpoints
type EncodingOutputsFtpAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/ftp/{output_id}/customData' endpoints
	Customdata *EncodingOutputsFtpCustomdataAPI
}

// NewEncodingOutputsFtpAPI constructor for EncodingOutputsFtpAPI that takes options as argument
func NewEncodingOutputsFtpAPI(options ...apiclient.APIClientOption) (*EncodingOutputsFtpAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsFtpAPIWithClient(apiClient), nil
}

// NewEncodingOutputsFtpAPIWithClient constructor for EncodingOutputsFtpAPI that takes an APIClient as argument
func NewEncodingOutputsFtpAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsFtpAPI {
	a := &EncodingOutputsFtpAPI{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsFtpCustomdataAPIWithClient(apiClient)

	return a
}

// Create FTP Output
func (api *EncodingOutputsFtpAPI) Create(ftpOutput model.FtpOutput) (*model.FtpOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.FtpOutput
	err := api.apiClient.Post("/encoding/outputs/ftp", &ftpOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete FTP Output
func (api *EncodingOutputsFtpAPI) Delete(outputId string) (*model.FtpOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.FtpOutput
	err := api.apiClient.Delete("/encoding/outputs/ftp/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get FTP Output Details
func (api *EncodingOutputsFtpAPI) Get(outputId string) (*model.FtpOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.FtpOutput
	err := api.apiClient.Get("/encoding/outputs/ftp/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List FTP Outputs
func (api *EncodingOutputsFtpAPI) List(queryParams ...func(*EncodingOutputsFtpAPIListQueryParams)) (*pagination.FtpOutputsListPagination, error) {
	queryParameters := &EncodingOutputsFtpAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.FtpOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/ftp", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsFtpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsFtpAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsFtpAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
