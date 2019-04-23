package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsExtractWebvttApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsExtractWebvttCustomdataApi
}

func NewEncodingEncodingsCaptionsExtractWebvttApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsExtractWebvttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsExtractWebvttApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsExtractWebvttCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsExtractWebvttApi) Get(encodingId string, captionsId string) (*model.WebVttExtract, error) {
    var resp *model.WebVttExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/webvtt/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractWebvttApi) List(encodingId string, queryParams ...func(*query.WebVttExtractListQueryParams)) (*pagination.WebVttExtractsListPagination, error) {
    queryParameters := &query.WebVttExtractListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.WebVttExtractsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/extract/webvtt", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractWebvttApi) Delete(encodingId string, captionsId string) (*model.WebVttExtract, error) {
    var resp *model.WebVttExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/extract/webvtt/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsExtractWebvttApi) Create(encodingId string, webVttExtract model.WebVttExtract) (*model.WebVttExtract, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.WebVttExtract(webVttExtract)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/extract/webvtt", &payload, reqParams)
    return &payload, err
}
