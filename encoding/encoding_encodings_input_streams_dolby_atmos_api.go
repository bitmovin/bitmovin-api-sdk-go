package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsDolbyAtmosApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsDolbyAtmosApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsDolbyAtmosApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsDolbyAtmosApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsDolbyAtmosApi) Create(encodingId string, dolbyAtmosIngestInputStream model.DolbyAtmosIngestInputStream) (*model.DolbyAtmosIngestInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.DolbyAtmosIngestInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos", &dolbyAtmosIngestInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsDolbyAtmosApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsDolbyAtmosApi) Get(encodingId string, inputStreamId string) (*model.DolbyAtmosIngestInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.DolbyAtmosIngestInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsDolbyAtmosApi) List(encodingId string, queryParams ...func(*query.DolbyAtmosIngestInputStreamListQueryParams)) (*pagination.DolbyAtmosIngestInputStreamsListPagination, error) {
    queryParameters := &query.DolbyAtmosIngestInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DolbyAtmosIngestInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos", nil, &responseModel, reqParams)
    return responseModel, err
}

