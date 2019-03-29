package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsTsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsTsCustomdataApi
    Drm *EncodingEncodingsMuxingsTsDrmApi
    Captions *EncodingEncodingsMuxingsTsCaptionsApi
}

func NewEncodingEncodingsMuxingsTsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsTsCustomdataApi(configs...)
    api.Customdata = customdataApi
    drmApi, err := NewEncodingEncodingsMuxingsTsDrmApi(configs...)
    api.Drm = drmApi
    captionsApi, err := NewEncodingEncodingsMuxingsTsCaptionsApi(configs...)
    api.Captions = captionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsTsApi) Get(encodingId string, muxingId string) (*model.TsMuxing, error) {
    var resp *model.TsMuxing
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsTsApi) Create(encodingId string, tsMuxing model.TsMuxing) (*model.TsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.TsMuxing(tsMuxing)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsTsApi) List(encodingId string, queryParams ...func(*query.TsMuxingListQueryParams)) (*pagination.TsMuxingsListPagination, error) {
    queryParameters := &query.TsMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TsMuxingsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts", &resp, reqParams)
    return resp, err
}
