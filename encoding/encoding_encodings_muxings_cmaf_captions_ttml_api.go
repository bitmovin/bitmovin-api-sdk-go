package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafCaptionsTtmlApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafCaptionsTtmlApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafCaptionsTtmlApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafCaptionsTtmlApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafCaptionsTtmlCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlApi) List(encodingId string, muxingId string, queryParams ...func(*query.TtmlEmbedListQueryParams)) (*pagination.TtmlEmbedsListPagination, error) {
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
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlApi) Delete(encodingId string, muxingId string, captionsId string) (*model.TtmlEmbed, error) {
    var resp *model.TtmlEmbed
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlApi) Get(encodingId string, muxingId string, captionsId string) (*model.TtmlEmbed, error) {
    var resp *model.TtmlEmbed
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlApi) Create(encodingId string, muxingId string, ttmlEmbed model.TtmlEmbed) (*model.TtmlEmbed, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.TtmlEmbed(ttmlEmbed)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml", &payload, reqParams)
    return &payload, err
}
