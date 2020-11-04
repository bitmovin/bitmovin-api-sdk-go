package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsSftpCustomdataAPI communicates with '/encoding/outputs/sftp/{output_id}/customData' endpoints
type EncodingOutputsSftpCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingOutputsSftpCustomdataAPI constructor for EncodingOutputsSftpCustomdataAPI that takes options as argument
func NewEncodingOutputsSftpCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsSftpCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsSftpCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsSftpCustomdataAPIWithClient constructor for EncodingOutputsSftpCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsSftpCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsSftpCustomdataAPI {
	a := &EncodingOutputsSftpCustomdataAPI{apiClient: apiClient}
	return a
}

// Get SFTP Output Custom Data
func (api *EncodingOutputsSftpCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/outputs/sftp/{output_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
