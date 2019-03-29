package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsSegmentedRawApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsSegmentedRawCustomdataApi
}

func NewEncodingEncodingsMuxingsSegmentedRawApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsSegmentedRawApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsSegmentedRawApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsSegmentedRawCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsSegmentedRawApi) List(encodingId string, queryParams ...func(*query.SegmentedRawMuxingListQueryParams)) (*pagination.SegmentedRawMuxingsListPagination, error) {
    queryParameters := &query.SegmentedRawMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SegmentedRawMuxingsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsSegmentedRawApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsSegmentedRawApi) Create(encodingId string, segmentedRawMuxing model.SegmentedRawMuxing) (*model.SegmentedRawMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.SegmentedRawMuxing(segmentedRawMuxing)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/segmented-raw", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsSegmentedRawApi) Get(encodingId string, muxingId string) (*model.SegmentedRawMuxing, error) {
    var resp *model.SegmentedRawMuxing
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}", &resp, reqParams)
    return resp, err
}
