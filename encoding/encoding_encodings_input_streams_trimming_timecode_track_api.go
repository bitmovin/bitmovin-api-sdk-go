package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi) List(encodingId string, queryParams ...func(*query.TimecodeTrackTrimmingInputStreamListQueryParams)) (*pagination.TimecodeTrackTrimmingInputStreamsListPagination, error) {
    queryParameters := &query.TimecodeTrackTrimmingInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TimecodeTrackTrimmingInputStreamsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi) Get(encodingId string, inputStreamId string) (*model.TimecodeTrackTrimmingInputStream, error) {
    var resp *model.TimecodeTrackTrimmingInputStream
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track/{input_stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi) Create(encodingId string, timecodeTrackTrimmingInputStream model.TimecodeTrackTrimmingInputStream) (*model.TimecodeTrackTrimmingInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.TimecodeTrackTrimmingInputStream(timecodeTrackTrimmingInputStream)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track/{input_stream_id}", &resp, reqParams)
    return resp, err
}
