package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsConcatenationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsConcatenationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsConcatenationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsConcatenationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsConcatenationApi) Get(encodingId string, inputStreamId string) (*model.ConcatenationInputStream, error) {
    var resp *model.ConcatenationInputStream
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/concatenation/{input_stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsConcatenationApi) List(encodingId string, queryParams ...func(*query.ConcatenationInputStreamListQueryParams)) (*pagination.ConcatenationInputStreamsListPagination, error) {
    queryParameters := &query.ConcatenationInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.ConcatenationInputStreamsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/concatenation", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsConcatenationApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/concatenation/{input_stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsConcatenationApi) Create(encodingId string, concatenationInputStream model.ConcatenationInputStream) (*model.ConcatenationInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.ConcatenationInputStream(concatenationInputStream)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/concatenation", &payload, reqParams)
    return &payload, err
}
