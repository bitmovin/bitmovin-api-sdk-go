package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// PlayerTestingCodecCompatibilityAPI communicates with '/player/testing/codec-compatibility' endpoints
type PlayerTestingCodecCompatibilityAPI struct {
	apiClient *apiclient.APIClient
}

// NewPlayerTestingCodecCompatibilityAPI constructor for PlayerTestingCodecCompatibilityAPI that takes options as argument
func NewPlayerTestingCodecCompatibilityAPI(options ...apiclient.APIClientOption) (*PlayerTestingCodecCompatibilityAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewPlayerTestingCodecCompatibilityAPIWithClient(apiClient), nil
}

// NewPlayerTestingCodecCompatibilityAPIWithClient constructor for PlayerTestingCodecCompatibilityAPI that takes an APIClient as argument
func NewPlayerTestingCodecCompatibilityAPIWithClient(apiClient *apiclient.APIClient) *PlayerTestingCodecCompatibilityAPI {
	a := &PlayerTestingCodecCompatibilityAPI{apiClient: apiClient}
	return a
}

// Get Codec Compatibility Report
func (api *PlayerTestingCodecCompatibilityAPI) Get(queryParams ...func(*PlayerTestingCodecCompatibilityAPIGetQueryParams)) (*model.PccReport, error) {
	queryParameters := &PlayerTestingCodecCompatibilityAPIGetQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel model.PccReport
	err := api.apiClient.Get("/player/testing/codec-compatibility", nil, &responseModel, reqParams)
	return &responseModel, err
}

// PlayerTestingCodecCompatibilityAPIGetQueryParams contains all query parameters for the Get endpoint
type PlayerTestingCodecCompatibilityAPIGetQueryParams struct {
	IncludePrerelease bool   `query:"includePrerelease"`
	ReportedOnly      bool   `query:"reportedOnly"`
	HdrOnly           bool   `query:"hdrOnly"`
	Codec             string `query:"codec"`
	Device            string `query:"device"`
}

// Params will return a map of query parameters
func (q *PlayerTestingCodecCompatibilityAPIGetQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
