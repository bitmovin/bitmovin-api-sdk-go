package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingSimpleAPI intermediary API object with no endpoints
type EncodingSimpleAPI struct {
	apiClient *apiclient.APIClient

	// Jobs intermediary API object with no endpoints
	Jobs *EncodingSimpleJobsAPI
}

// NewEncodingSimpleAPI constructor for EncodingSimpleAPI that takes options as argument
func NewEncodingSimpleAPI(options ...apiclient.APIClientOption) (*EncodingSimpleAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingSimpleAPIWithClient(apiClient), nil
}

// NewEncodingSimpleAPIWithClient constructor for EncodingSimpleAPI that takes an APIClient as argument
func NewEncodingSimpleAPIWithClient(apiClient *apiclient.APIClient) *EncodingSimpleAPI {
	a := &EncodingSimpleAPI{apiClient: apiClient}
	a.Jobs = NewEncodingSimpleJobsAPIWithClient(apiClient)

	return a
}
