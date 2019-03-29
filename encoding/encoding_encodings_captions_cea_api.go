package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsCeaApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsCeaCustomdataApi
}

func NewEncodingEncodingsCaptionsCeaApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsCeaApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsCeaApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsCeaCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsCeaApi) Delete(encodingId string, captionsId string) (*model.CaptionsCea, error) {
    var resp *model.CaptionsCea
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/cea/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsCeaApi) Create(encodingId string, captionsCea model.CaptionsCea) (*model.CaptionsCea, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.CaptionsCea(captionsCea)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/cea", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsCaptionsCeaApi) Get(encodingId string, captionsId string) (*model.CaptionsCea, error) {
    var resp *model.CaptionsCea
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/cea/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsCeaApi) List(encodingId string, queryParams ...func(*query.CaptionsCeaListQueryParams)) (*pagination.CaptionsCeasListPagination, error) {
    queryParameters := &query.CaptionsCeaListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CaptionsCeasListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/cea", &resp, reqParams)
    return resp, err
}
