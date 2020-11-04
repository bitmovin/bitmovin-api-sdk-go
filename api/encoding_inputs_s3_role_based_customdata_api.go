package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsS3RoleBasedCustomdataAPI communicates with '/encoding/inputs/s3-role-based/{input_id}/customData' endpoints
type EncodingInputsS3RoleBasedCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsS3RoleBasedCustomdataAPI constructor for EncodingInputsS3RoleBasedCustomdataAPI that takes options as argument
func NewEncodingInputsS3RoleBasedCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsS3RoleBasedCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsS3RoleBasedCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsS3RoleBasedCustomdataAPIWithClient constructor for EncodingInputsS3RoleBasedCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsS3RoleBasedCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsS3RoleBasedCustomdataAPI {
	a := &EncodingInputsS3RoleBasedCustomdataAPI{apiClient: apiClient}
	return a
}

// Get S3 Custom Data
func (api *EncodingInputsS3RoleBasedCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/inputs/s3-role-based/{input_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
