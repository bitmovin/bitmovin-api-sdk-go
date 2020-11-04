package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsCaptionsSccCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/captions/scc/{captions_id}/customData' endpoints
type EncodingEncodingsCaptionsSccCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsCaptionsSccCustomdataAPI constructor for EncodingEncodingsCaptionsSccCustomdataAPI that takes options as argument
func NewEncodingEncodingsCaptionsSccCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsCaptionsSccCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsCaptionsSccCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsCaptionsSccCustomdataAPIWithClient constructor for EncodingEncodingsCaptionsSccCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsCaptionsSccCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsCaptionsSccCustomdataAPI {
	a := &EncodingEncodingsCaptionsSccCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Convert SCC captions Custom Data
func (api *EncodingEncodingsCaptionsSccCustomdataAPI) Get(encodingId string, captionsId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["captions_id"] = captionsId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/scc/{captions_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
