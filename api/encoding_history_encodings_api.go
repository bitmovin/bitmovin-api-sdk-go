package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingHistoryEncodingsAPI communicates with '/encoding/history/encodings' endpoints
type EncodingHistoryEncodingsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingHistoryEncodingsAPI constructor for EncodingHistoryEncodingsAPI that takes options as argument
func NewEncodingHistoryEncodingsAPI(options ...apiclient.APIClientOption) (*EncodingHistoryEncodingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingHistoryEncodingsAPIWithClient(apiClient), nil
}

// NewEncodingHistoryEncodingsAPIWithClient constructor for EncodingHistoryEncodingsAPI that takes an APIClient as argument
func NewEncodingHistoryEncodingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingHistoryEncodingsAPI {
	a := &EncodingHistoryEncodingsAPI{apiClient: apiClient}
	return a
}

// Get (Experimental) History Encoding Details
// Returns every configuration and result related to an Encoding.
func (api *EncodingHistoryEncodingsAPI) Get(encodingId string) (*model.HistoryEncoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.HistoryEncoding
	err := api.apiClient.Get("/encoding/history/encodings/{encoding_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List (Experimental) List all History Encodings
// Returns a list of History Encodings. Recently created Encodings might not be available as History Encoding.
func (api *EncodingHistoryEncodingsAPI) List(queryParams ...func(*EncodingHistoryEncodingsAPIListQueryParams)) (*pagination.EncodingsListPagination, error) {
	queryParameters := &EncodingHistoryEncodingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EncodingsListPagination
	err := api.apiClient.Get("/encoding/history/encodings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingHistoryEncodingsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingHistoryEncodingsAPIListQueryParams struct {
	Offset                 int32              `query:"offset"`
	Limit                  int32              `query:"limit"`
	Sort                   string             `query:"sort"`
	Type_                  string             `query:"type"`
	Status                 string             `query:"status"`
	CloudRegion            model.CloudRegion  `query:"cloudRegion"`
	SelectedCloudRegion    model.CloudRegion  `query:"selectedCloudRegion"`
	EncoderVersion         string             `query:"encoderVersion"`
	SelectedEncoderVersion string             `query:"selectedEncoderVersion"`
	SelectedEncodingMode   model.EncodingMode `query:"selectedEncodingMode"`
	Name                   string             `query:"name"`
	Search                 string             `query:"search"`
	CreatedAtNewerThan     model.DateTime     `query:"createdAtNewerThan"`
	CreatedAtOlderThan     model.DateTime     `query:"createdAtOlderThan"`
	StartedAtNewerThan     model.DateTime     `query:"startedAtNewerThan"`
	StartedAtOlderThan     model.DateTime     `query:"startedAtOlderThan"`
	FinishedAtNewerThan    model.DateTime     `query:"finishedAtNewerThan"`
	FinishedAtOlderThan    model.DateTime     `query:"finishedAtOlderThan"`
}

// Params will return a map of query parameters
func (q *EncodingHistoryEncodingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
