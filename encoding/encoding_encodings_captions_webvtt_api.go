package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsWebvttApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsWebvttCustomdataApi
}

func NewEncodingEncodingsCaptionsWebvttApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsWebvttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsWebvttApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsWebvttCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsWebvttApi) Create(encodingId string, webVttExtract model.WebVttExtract) (*model.WebVttExtract, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.WebVttExtract(webVttExtract)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/webvtt", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsCaptionsWebvttApi) List(encodingId string, queryParams ...func(*query.WebVttExtractListQueryParams)) (*pagination.WebVttExtractsListPagination, error) {
    queryParameters := &query.WebVttExtractListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.WebVttExtractsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/webvtt", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsWebvttApi) Delete(encodingId string, captionsId string) (*model.WebVttExtract, error) {
    var resp *model.WebVttExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/webvtt/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsWebvttApi) Get(encodingId string, captionsId string) (*model.WebVttExtract, error) {
    var resp *model.WebVttExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/webvtt/{captions_id}", &resp, reqParams)
    return resp, err
}
