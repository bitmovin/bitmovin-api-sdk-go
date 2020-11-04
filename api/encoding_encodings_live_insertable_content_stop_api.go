package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingEncodingsLiveInsertableContentStopAPI communicates with '/encoding/encodings/{encoding_id}/live/insertable-content/stop' endpoints
type EncodingEncodingsLiveInsertableContentStopAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveInsertableContentStopAPI constructor for EncodingEncodingsLiveInsertableContentStopAPI that takes options as argument
func NewEncodingEncodingsLiveInsertableContentStopAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveInsertableContentStopAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveInsertableContentStopAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveInsertableContentStopAPIWithClient constructor for EncodingEncodingsLiveInsertableContentStopAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveInsertableContentStopAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveInsertableContentStopAPI {
	a := &EncodingEncodingsLiveInsertableContentStopAPI{apiClient: apiClient}
	return a
}

// Create Stops Currently Running Inserted Content
func (api *EncodingEncodingsLiveInsertableContentStopAPI) Create(encodingId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/insertable-content/stop", nil, nil, reqParams)
	return err
}
