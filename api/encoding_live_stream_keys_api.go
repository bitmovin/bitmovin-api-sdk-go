package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingLiveStreamKeysAPI communicates with '/encoding/live/stream-keys' endpoints
type EncodingLiveStreamKeysAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingLiveStreamKeysAPI constructor for EncodingLiveStreamKeysAPI that takes options as argument
func NewEncodingLiveStreamKeysAPI(options ...apiclient.APIClientOption) (*EncodingLiveStreamKeysAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingLiveStreamKeysAPIWithClient(apiClient), nil
}

// NewEncodingLiveStreamKeysAPIWithClient constructor for EncodingLiveStreamKeysAPI that takes an APIClient as argument
func NewEncodingLiveStreamKeysAPIWithClient(apiClient *apiclient.APIClient) *EncodingLiveStreamKeysAPI {
	a := &EncodingLiveStreamKeysAPI{apiClient: apiClient}
	return a
}

// Create new stream key
func (api *EncodingLiveStreamKeysAPI) Create(streamKey model.StreamKey) (*model.StreamKey, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamKey
	err := api.apiClient.Post("/encoding/live/stream-keys", &streamKey, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Stream Key
func (api *EncodingLiveStreamKeysAPI) Delete(streamKeyId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_key_id"] = streamKeyId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/live/stream-keys/{stream_key_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Stream Key details
func (api *EncodingLiveStreamKeysAPI) Get(streamKeyId string) (*model.StreamKey, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_key_id"] = streamKeyId
	}

	var responseModel model.StreamKey
	err := api.apiClient.Get("/encoding/live/stream-keys/{stream_key_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Stream Keys
func (api *EncodingLiveStreamKeysAPI) List(queryParams ...func(*EncodingLiveStreamKeysAPIListQueryParams)) (*pagination.StreamKeysListPagination, error) {
	queryParameters := &EncodingLiveStreamKeysAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamKeysListPagination
	err := api.apiClient.Get("/encoding/live/stream-keys", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingLiveStreamKeysAPIListQueryParams contains all query parameters for the List endpoint
type EncodingLiveStreamKeysAPIListQueryParams struct {
	Offset                int32  `query:"offset"`
	Limit                 int32  `query:"limit"`
	Sort                  string `query:"sort"`
	Type_                 string `query:"type"`
	Status                string `query:"status"`
	AssignedIngestPointId string `query:"assignedIngestPointId"`
	AssignedEncodingId    string `query:"assignedEncodingId"`
}

// Params will return a map of query parameters
func (q *EncodingLiveStreamKeysAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
