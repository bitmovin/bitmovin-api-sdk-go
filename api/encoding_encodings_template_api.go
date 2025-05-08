package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsTemplateAPI communicates with '/encoding/encodings/{encoding_id}/template' endpoints
type EncodingEncodingsTemplateAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsTemplateAPI constructor for EncodingEncodingsTemplateAPI that takes options as argument
func NewEncodingEncodingsTemplateAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsTemplateAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsTemplateAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsTemplateAPIWithClient constructor for EncodingEncodingsTemplateAPI that takes an APIClient as argument
func NewEncodingEncodingsTemplateAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsTemplateAPI {
	a := &EncodingEncodingsTemplateAPI{apiClient: apiClient}
	return a
}

// Get Encoding Template URL
// Retrieve the download URL for the Encoding Template associated with a specific encoding.
func (api *EncodingEncodingsTemplateAPI) Get(encodingId string) (*model.EncodingTemplateUrlResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.EncodingTemplateUrlResponse
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/template", nil, &responseModel, reqParams)
	return &responseModel, err
}
