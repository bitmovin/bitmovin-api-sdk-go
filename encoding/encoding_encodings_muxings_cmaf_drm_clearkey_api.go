package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafDrmClearkeyApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafDrmClearkeyCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafDrmClearkeyApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmClearkeyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmClearkeyApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafDrmClearkeyCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmClearkeyApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/clearkey/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmClearkeyApi) Get(encodingId string, muxingId string, drmId string) (*model.ClearKeyDrm, error) {
    var resp *model.ClearKeyDrm
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/clearkey/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmClearkeyApi) List(encodingId string, muxingId string, queryParams ...func(*query.ClearKeyDrmListQueryParams)) (*pagination.ClearKeyDrmsListPagination, error) {
    queryParameters := &query.ClearKeyDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.ClearKeyDrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/clearkey", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmClearkeyApi) Create(encodingId string, muxingId string, clearKeyDrm model.ClearKeyDrm) (*model.ClearKeyDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.ClearKeyDrm(clearKeyDrm)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/clearkey", &payload, reqParams)
    return &payload, err
}