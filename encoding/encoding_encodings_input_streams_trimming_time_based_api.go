package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsTrimmingTimeBasedApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsTrimmingTimeBasedApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsTrimmingTimeBasedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsTrimmingTimeBasedApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedApi) Create(encodingId string, timeBasedTrimmingInputStream model.TimeBasedTrimmingInputStream) (*model.TimeBasedTrimmingInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.TimeBasedTrimmingInputStream(timeBasedTrimmingInputStream)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based/{input_stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedApi) List(encodingId string, queryParams ...func(*query.TimeBasedTrimmingInputStreamListQueryParams)) (*pagination.TimeBasedTrimmingInputStreamsListPagination, error) {
    queryParameters := &query.TimeBasedTrimmingInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TimeBasedTrimmingInputStreamsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedApi) Get(encodingId string, inputStreamId string) (*model.TimeBasedTrimmingInputStream, error) {
    var resp *model.TimeBasedTrimmingInputStream
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based/{input_stream_id}", &resp, reqParams)
    return resp, err
}
