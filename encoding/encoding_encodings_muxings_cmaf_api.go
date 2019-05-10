package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafCustomdataApi
    Captions *EncodingEncodingsMuxingsCmafCaptionsApi
}

func NewEncodingEncodingsMuxingsCmafApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafCustomdataApi(configs...)
    api.Customdata = customdataApi
    captionsApi, err := NewEncodingEncodingsMuxingsCmafCaptionsApi(configs...)
    api.Captions = captionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafApi) List(encodingId string, queryParams ...func(*query.CmafMuxingListQueryParams)) (*pagination.CmafMuxingsListPagination, error) {
    queryParameters := &query.CmafMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CmafMuxingsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafApi) Get(encodingId string, muxingId string) (*model.CmafMuxing, error) {
    var resp *model.CmafMuxing
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafApi) Create(encodingId string, cmafMuxing model.CmafMuxing) (*model.CmafMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.CmafMuxing(cmafMuxing)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf", &payload, reqParams)
    return &payload, err
}
