package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingTemplatesAPI communicates with '/encoding/templates/start' endpoints
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

// Start BETA: Start an Encoding defined with an Encoding Template
func (api *EncodingTemplatesAPI) Start(encodingTemplateRequest interface{}) (*model.EncodingTemplateStartResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.EncodingTemplateStartResponse
	err := api.apiClient.Post("/encoding/templates/start", &encodingTemplateRequest, &responseModel, reqParams)
	return &responseModel, err
}
