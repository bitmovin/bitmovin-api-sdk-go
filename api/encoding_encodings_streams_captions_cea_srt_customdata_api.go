package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt/{captions_id}/customData' endpoints
type EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI constructor for EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPIWithClient constructor for EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI {
	a := &EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Embed SRT captions as 608/708 Custom Data
func (api *EncodingEncodingsStreamsCaptionsCeaSrtCustomdataAPI) Get(encodingId string, streamId string, captionsId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["captions_id"] = captionsId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/srt/{captions_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
