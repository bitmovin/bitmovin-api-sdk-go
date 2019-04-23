package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsSidecarsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsSidecarsCustomdataApi
    Webvtt *EncodingEncodingsSidecarsWebvttApi
}

func NewEncodingEncodingsSidecarsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSidecarsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSidecarsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsSidecarsCustomdataApi(configs...)
    api.Customdata = customdataApi
    webvttApi, err := NewEncodingEncodingsSidecarsWebvttApi(configs...)
    api.Webvtt = webvttApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsSidecarsApi) Delete(encodingId string, sidecarId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsSidecarsApi) List(encodingId string, queryParams ...func(*query.SidecarFileListQueryParams)) (*pagination.SidecarFilesListPagination, error) {
    queryParameters := &query.SidecarFileListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SidecarFilesListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsSidecarsApi) Create(encodingId string, sidecarFile model.SidecarFile) (*model.SidecarFile, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.SidecarFile(sidecarFile)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/sidecars", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsSidecarsApi) Get(encodingId string, sidecarId string) (*model.SidecarFile, error) {
    var resp *model.SidecarFile
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}", &resp, reqParams)
    return resp, err
}
