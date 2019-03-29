package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsFiltersApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsFiltersApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsFiltersApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsFiltersApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsFiltersApi) List(encodingId string, streamId string, queryParams ...func(*query.StreamFilterListListQueryParams)) (*model.StreamFilterList, error) {
    queryParameters := &query.StreamFilterListListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *model.StreamFilterList
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsFiltersApi) DeleteAll(encodingId string, streamId string) (*model.BitmovinResponseList, error) {
    var resp *model.BitmovinResponseList
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsFiltersApi) Delete(encodingId string, streamId string, filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsStreamsFiltersApi) Create(encodingId string, streamId string, streamFilter model.StreamFilterList) (*model.StreamFilterList, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }
    payload := model.StreamFilterList(streamFilter)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", &payload, reqParams)
    return &payload, err
}
