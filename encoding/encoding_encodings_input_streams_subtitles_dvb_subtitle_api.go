package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi) Create(encodingId string, dvbSubtitleInputStream model.DvbSubtitleInputStream) (*model.DvbSubtitleInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.DvbSubtitleInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle", &dvbSubtitleInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi) Get(encodingId string, inputStreamId string) (*model.DvbSubtitleInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.DvbSubtitleInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleApi) List(encodingId string, queryParams ...func(*query.DvbSubtitleInputStreamListQueryParams)) (*pagination.DvbSubtitleInputStreamsListPagination, error) {
    queryParameters := &query.DvbSubtitleInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DvbSubtitleInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle", nil, &responseModel, reqParams)
    return responseModel, err
}

