package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingStatisticsCdnAPI intermediary API object with no endpoints
type EncodingStatisticsCdnAPI struct {
	apiClient *apiclient.APIClient

	// Usage communicates with '/encoding/statistics/cdn/usage' endpoints
	Usage *EncodingStatisticsCdnUsageAPI
}

// NewEncodingStatisticsCdnAPI constructor for EncodingStatisticsCdnAPI that takes options as argument
func NewEncodingStatisticsCdnAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsCdnAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsCdnAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsCdnAPIWithClient constructor for EncodingStatisticsCdnAPI that takes an APIClient as argument
func NewEncodingStatisticsCdnAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsCdnAPI {
	a := &EncodingStatisticsCdnAPI{apiClient: apiClient}
	a.Usage = NewEncodingStatisticsCdnUsageAPIWithClient(apiClient)

	return a
}
