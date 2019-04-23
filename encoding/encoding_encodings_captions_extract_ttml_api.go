package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsExtractTtmlApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsExtractTtmlCustomdataApi
}

func NewEncodingEncodingsCaptionsExtractTtmlApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsExtractTtmlApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsExtractTtmlApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsExtractTtmlCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsExtractTtmlApi) Delete(encodingId string, captionsId string) (*model.TtmlExtract, error) {
    var resp *model.TtmlExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/extract/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractTtmlApi) Create(encodingId string, ttmlExtract model.TtmlExtract) (*model.TtmlExtract, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.TtmlExtract(ttmlExtract)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/extract/ttml", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsCaptionsExtractTtmlApi) List(encodingId string, queryParams ...func(*query.TtmlExtractListQueryParams)) (*pagination.TtmlExtractsListPagination, error) {
    queryParameters := &query.TtmlExtractListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TtmlExtractsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/ttml", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractTtmlApi) Get(encodingId string, captionsId string) (*model.TtmlExtract, error) {
    var resp *model.TtmlExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
