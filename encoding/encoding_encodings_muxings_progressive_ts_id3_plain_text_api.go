package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataApi
}

func NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi) List(encodingId string, muxingId string, queryParams ...func(*query.PlaintextId3TagListQueryParams)) (*pagination.PlaintextId3TagsListPagination, error) {
    queryParameters := &query.PlaintextId3TagListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.PlaintextId3TagsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi) Delete(encodingId string, muxingId string, id3TagId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi) Create(encodingId string, muxingId string, plaintextId3Tag model.PlaintextId3Tag) (*model.PlaintextId3Tag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.PlaintextId3Tag(plaintextId3Tag)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi) Get(encodingId string, muxingId string, id3TagId string) (*model.PlaintextId3Tag, error) {
    var resp *model.PlaintextId3Tag
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}", &resp, reqParams)
    return resp, err
}
