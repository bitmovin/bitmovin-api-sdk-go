package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4CaptionsTtmlApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4CaptionsTtmlCustomdataApi
}

func NewEncodingEncodingsMuxingsFmp4CaptionsTtmlApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4CaptionsTtmlApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4CaptionsTtmlApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4CaptionsTtmlCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4CaptionsTtmlApi) Get(encodingId string, muxingId string, captionsId string) (*model.TtmlEmbed, error) {
    var resp *model.TtmlEmbed
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4CaptionsTtmlApi) Create(encodingId string, muxingId string, ttmlEmbed model.TtmlEmbed) (*model.TtmlEmbed, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.TtmlEmbed(ttmlEmbed)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/ttml", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsFmp4CaptionsTtmlApi) List(encodingId string, muxingId string, queryParams ...func(*query.TtmlEmbedListQueryParams)) (*pagination.TtmlEmbedsListPagination, error) {
    queryParameters := &query.TtmlEmbedListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TtmlEmbedsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/ttml", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsFmp4CaptionsTtmlApi) Delete(encodingId string, muxingId string, captionsId string) (*model.TtmlEmbed, error) {
    var resp *model.TtmlEmbed
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/captions/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
