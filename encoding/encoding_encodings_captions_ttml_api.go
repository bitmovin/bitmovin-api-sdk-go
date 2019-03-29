package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsTtmlApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsTtmlCustomdataApi
}

func NewEncodingEncodingsCaptionsTtmlApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsTtmlApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsTtmlApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsTtmlCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsTtmlApi) List(encodingId string, queryParams ...func(*query.TtmlExtractListQueryParams)) (*pagination.TtmlExtractsListPagination, error) {
    queryParameters := &query.TtmlExtractListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.TtmlExtractsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/ttml", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsTtmlApi) Get(encodingId string, captionsId string) (*model.TtmlExtract, error) {
    var resp *model.TtmlExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsCaptionsTtmlApi) Create(encodingId string, ttmlExtract model.TtmlExtract) (*model.TtmlExtract, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.TtmlExtract(ttmlExtract)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/ttml", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsCaptionsTtmlApi) Delete(encodingId string, captionsId string) (*model.TtmlExtract, error) {
    var resp *model.TtmlExtract
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/ttml/{captions_id}", &resp, reqParams)
    return resp, err
}
