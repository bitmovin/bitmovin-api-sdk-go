package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMxfApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMxfCustomdataApi
}

func NewEncodingEncodingsMuxingsMxfApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMxfApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMxfApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMxfCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMxfApi) Create(encodingId string, mxfMuxing model.MxfMuxing) (*model.MxfMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.MxfMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mxf", &mxfMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMxfApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMxfApi) Get(encodingId string, muxingId string) (*model.MxfMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.MxfMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mxf/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMxfApi) List(encodingId string, queryParams ...func(*query.MxfMuxingListQueryParams)) (*pagination.MxfMuxingsListPagination, error) {
    queryParameters := &query.MxfMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.MxfMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mxf", nil, &responseModel, reqParams)
    return responseModel, err
}

