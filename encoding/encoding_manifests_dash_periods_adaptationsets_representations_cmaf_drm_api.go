package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi struct {
    apiClient *common.ApiClient
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmContentprotectionApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi{apiClient: apiClient}

    contentprotectionApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmContentprotectionApi(configs...)
    api.Contentprotection = contentprotectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*query.DashCmafDrmRepresentationListQueryParams)) (*pagination.DashCmafDrmRepresentationsListPagination, error) {
    queryParameters := &query.DashCmafDrmRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.DashCmafDrmRepresentationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/drm", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
	}
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/drm/{representation_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi) Create(manifestId string, periodId string, adaptationsetId string, dashCmafDrmRepresentation model.DashCmafDrmRepresentation) (*model.DashCmafDrmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }
    payload := model.DashCmafDrmRepresentation(dashCmafDrmRepresentation)
    
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/drm", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashCmafDrmRepresentation, error) {
    var resp *model.DashCmafDrmRepresentation
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/drm/{representation_id}", &resp, reqParams)
    return resp, err
}
