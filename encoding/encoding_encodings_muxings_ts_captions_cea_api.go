package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsTsCaptionsCeaApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsTsCaptionsCeaCustomdataApi
}

func NewEncodingEncodingsMuxingsTsCaptionsCeaApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsCaptionsCeaApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsCaptionsCeaApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsTsCaptionsCeaCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsCaptionsCeaApi) List(encodingId string, muxingId string, queryParams ...func(*query.CaptionsEmbedCeaListQueryParams)) (*pagination.CaptionsEmbedCeasListPagination, error) {
    queryParameters := &query.CaptionsEmbedCeaListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CaptionsEmbedCeasListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/captions/cea", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsTsCaptionsCeaApi) Create(encodingId string, muxingId string, captionsEmbedCea model.CaptionsEmbedCea) (*model.CaptionsEmbedCea, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.CaptionsEmbedCea(captionsEmbedCea)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/captions/cea", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsTsCaptionsCeaApi) Delete(encodingId string, muxingId string, captionsId string) (*model.CaptionsEmbedCea, error) {
    var resp *model.CaptionsEmbedCea
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/captions/cea/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsTsCaptionsCeaApi) Get(encodingId string, muxingId string, captionsId string) (*model.CaptionsEmbedCea, error) {
    var resp *model.CaptionsEmbedCea
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/captions/cea/{captions_id}", &resp, reqParams)
    return resp, err
}
