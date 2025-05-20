package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingTemplatesAPI communicates with '/encoding/templates' endpoints
type EncodingTemplatesAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingTemplatesAPI constructor for EncodingTemplatesAPI that takes options as argument
func NewEncodingTemplatesAPI(options ...apiclient.APIClientOption) (*EncodingTemplatesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingTemplatesAPIWithClient(apiClient), nil
}

// NewEncodingTemplatesAPIWithClient constructor for EncodingTemplatesAPI that takes an APIClient as argument
func NewEncodingTemplatesAPIWithClient(apiClient *apiclient.APIClient) *EncodingTemplatesAPI {
	a := &EncodingTemplatesAPI{apiClient: apiClient}
	return a
}

// Create Store an Encoding Template
// Stores an Encoding Template that can be used later to start encodings or standby pools
func (api *EncodingTemplatesAPI) Create(encodingTemplateRequest interface{}) (*model.EncodingTemplateDetails, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.EncodingTemplateDetails
	err := api.apiClient.Post("/encoding/templates", &encodingTemplateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Encoding Template
func (api *EncodingTemplatesAPI) Delete(encodingTemplateId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_template_id"] = encodingTemplateId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/templates/{encoding_template_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Encoding Template details
// Get the details of a stored Encoding Template
func (api *EncodingTemplatesAPI) Get(encodingTemplateId string) (*model.EncodingTemplateDetails, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_template_id"] = encodingTemplateId
	}

	var responseModel model.EncodingTemplateDetails
	err := api.apiClient.Get("/encoding/templates/{encoding_template_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List stored Encoding Templates
func (api *EncodingTemplatesAPI) List(queryParams ...func(*EncodingTemplatesAPIListQueryParams)) (*pagination.EncodingTemplateResponsesListPagination, error) {
	queryParameters := &EncodingTemplatesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EncodingTemplateResponsesListPagination
	err := api.apiClient.Get("/encoding/templates", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start an Encoding defined with an Encoding Template
func (api *EncodingTemplatesAPI) Start(encodingTemplateRequest interface{}) (*model.EncodingTemplateStartResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.EncodingTemplateStartResponse
	err := api.apiClient.Post("/encoding/templates/start", &encodingTemplateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingTemplatesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingTemplatesAPIListQueryParams struct {
	Offset int32                      `query:"offset"`
	Limit  int32                      `query:"limit"`
	Type_  model.EncodingTemplateType `query:"type"`
}

// Params will return a map of query parameters
func (q *EncodingTemplatesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
