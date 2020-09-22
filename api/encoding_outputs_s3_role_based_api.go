package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsS3RoleBasedAPI communicates with '/encoding/outputs/s3-role-based' endpoints
type EncodingOutputsS3RoleBasedAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/outputs/s3-role-based/{output_id}/customData' endpoints
	Customdata *EncodingOutputsS3RoleBasedCustomdataAPI
}

// NewEncodingOutputsS3RoleBasedAPI constructor for EncodingOutputsS3RoleBasedAPI that takes options as argument
func NewEncodingOutputsS3RoleBasedAPI(options ...apiclient.APIClientOption) (*EncodingOutputsS3RoleBasedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsS3RoleBasedAPIWithClient(apiClient), nil
}

// NewEncodingOutputsS3RoleBasedAPIWithClient constructor for EncodingOutputsS3RoleBasedAPI that takes an APIClient as argument
func NewEncodingOutputsS3RoleBasedAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsS3RoleBasedAPI {
	a := &EncodingOutputsS3RoleBasedAPI{apiClient: apiClient}
	a.Customdata = NewEncodingOutputsS3RoleBasedCustomdataAPIWithClient(apiClient)

	return a
}

// Create S3 Role-based Output
func (api *EncodingOutputsS3RoleBasedAPI) Create(s3RoleBasedOutput model.S3RoleBasedOutput) (*model.S3RoleBasedOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.S3RoleBasedOutput
	err := api.apiClient.Post("/encoding/outputs/s3-role-based", &s3RoleBasedOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete S3 Role-based Output
func (api *EncodingOutputsS3RoleBasedAPI) Delete(outputId string) (*model.S3RoleBasedOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.S3RoleBasedOutput
	err := api.apiClient.Delete("/encoding/outputs/s3-role-based/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get S3 Role-based Output Details
func (api *EncodingOutputsS3RoleBasedAPI) Get(outputId string) (*model.S3RoleBasedOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.S3RoleBasedOutput
	err := api.apiClient.Get("/encoding/outputs/s3-role-based/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List S3 Role-based Outputs
func (api *EncodingOutputsS3RoleBasedAPI) List(queryParams ...func(*EncodingOutputsS3RoleBasedAPIListQueryParams)) (*pagination.S3RoleBasedOutputsListPagination, error) {
	queryParameters := &EncodingOutputsS3RoleBasedAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.S3RoleBasedOutputsListPagination
	err := api.apiClient.Get("/encoding/outputs/s3-role-based", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsS3RoleBasedAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsS3RoleBasedAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsS3RoleBasedAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
