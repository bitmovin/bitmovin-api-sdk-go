package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsExtractCeaApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsExtractCeaCustomdataApi
}

func NewEncodingEncodingsCaptionsExtractCeaApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsExtractCeaApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsExtractCeaApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsExtractCeaCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsExtractCeaApi) List(encodingId string, queryParams ...func(*query.CaptionsCeaListQueryParams)) (*pagination.CaptionsCeasListPagination, error) {
    queryParameters := &query.CaptionsCeaListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CaptionsCeasListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/cea", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractCeaApi) Get(encodingId string, captionsId string) (*model.CaptionsCea, error) {
    var resp *model.CaptionsCea
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/cea/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractCeaApi) Delete(encodingId string, captionsId string) (*model.CaptionsCea, error) {
    var resp *model.CaptionsCea
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/extract/cea/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractCeaApi) Create(encodingId string, captionsCea model.CaptionsCea) (*model.CaptionsCea, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.CaptionsCea(captionsCea)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/extract/cea", &payload, reqParams)
    return &payload, err
}
