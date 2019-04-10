package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafDrmMarlinApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafDrmMarlinApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmMarlinApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmMarlinApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafDrmMarlinCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmMarlinApi) Create(encodingId string, muxingId string, marlinDrm model.MarlinDrm) (*model.MarlinDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.MarlinDrm(marlinDrm)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/marlin", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsCmafDrmMarlinApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/marlin/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmMarlinApi) Get(encodingId string, muxingId string, drmId string) (*model.MarlinDrm, error) {
    var resp *model.MarlinDrm
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/marlin/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmMarlinApi) List(encodingId string, muxingId string, queryParams ...func(*query.MarlinDrmListQueryParams)) (*pagination.MarlinDrmsListPagination, error) {
    queryParameters := &query.MarlinDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.MarlinDrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/marlin", &resp, reqParams)
    return resp, err
}
