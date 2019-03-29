package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsIngestApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsIngestApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsIngestApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsIngestApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsIngestApi) Create(encodingId string, ingestInputStream model.IngestInputStream) (*model.IngestInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.IngestInputStream(ingestInputStream)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/ingest", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsInputStreamsIngestApi) Get(encodingId string, inputStreamId string) (*model.IngestInputStream, error) {
    var resp *model.IngestInputStream
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/ingest/{input_stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsIngestApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/ingest/{input_stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsInputStreamsIngestApi) List(encodingId string, queryParams ...func(*query.IngestInputStreamListQueryParams)) (*pagination.IngestInputStreamsListPagination, error) {
    queryParameters := &query.IngestInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.IngestInputStreamsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/ingest", &resp, reqParams)
    return resp, err
}
