package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsSmoothContentprotectionApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsSmoothContentprotectionApi(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothContentprotectionApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothContentprotectionApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothContentprotectionApi) Delete(manifestId string, protectionId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["protection_id"] = protectionId
	}
    err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}/contentprotection/{protection_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsSmoothContentprotectionApi) List(manifestId string, queryParams ...func(*query.SmoothManifestContentProtectionListQueryParams)) (*pagination.SmoothManifestContentProtectionsListPagination, error) {
    queryParameters := &query.SmoothManifestContentProtectionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SmoothManifestContentProtectionsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/contentprotection", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsSmoothContentprotectionApi) Create(manifestId string, smoothManifestContentProtection model.SmoothManifestContentProtection) (*model.SmoothManifestContentProtection, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.SmoothManifestContentProtection(smoothManifestContentProtection)
    
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/contentprotection", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsSmoothContentprotectionApi) Get(manifestId string, protectionId string) (*model.SmoothManifestContentProtection, error) {
    var resp *model.SmoothManifestContentProtection
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["protection_id"] = protectionId
	}
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/contentprotection/{protection_id}", &resp, reqParams)
    return resp, err
}
