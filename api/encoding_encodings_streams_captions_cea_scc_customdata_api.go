package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}/customData' endpoints
type EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI constructor for EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI that takes options as argument
func NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataAPIWithClient constructor for EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsCaptionsCeaSccCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI {
	a := &EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI{apiClient: apiClient}
	return a
}

// Get Embed SCC captions as 608/708 Custom Data
func (api *EncodingEncodingsStreamsCaptionsCeaSccCustomdataAPI) Get(encodingId string, streamId string, captionsId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["captions_id"] = captionsId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/captions/608-708/scc/{captions_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
