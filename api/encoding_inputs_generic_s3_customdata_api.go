package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsGenericS3CustomdataAPI communicates with '/encoding/inputs/generic-s3/{input_id}/customData' endpoints
type EncodingInputsGenericS3CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsGenericS3CustomdataAPI constructor for EncodingInputsGenericS3CustomdataAPI that takes options as argument
func NewEncodingInputsGenericS3CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsGenericS3CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsGenericS3CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsGenericS3CustomdataAPIWithClient constructor for EncodingInputsGenericS3CustomdataAPI that takes an APIClient as argument
func NewEncodingInputsGenericS3CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsGenericS3CustomdataAPI {
	a := &EncodingInputsGenericS3CustomdataAPI{apiClient: apiClient}
	return a
}

// Get Generic S3 Input Custom Data
func (api *EncodingInputsGenericS3CustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/generic-s3/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
