package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsCustomXmlElementsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashPeriodsCustomXmlElementsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsCustomXmlElementsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsCustomXmlElementsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) List(manifestId string, periodId string, queryParams ...func(*query.CustomXmlElementListQueryParams)) (*pagination.CustomXmlElementsListPagination, error) {
    queryParameters := &query.CustomXmlElementListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CustomXmlElementsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) Create(manifestId string, periodId string, customXmlElement model.CustomXmlElement) (*model.CustomXmlElement, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }
    payload := model.CustomXmlElement(customXmlElement)
    
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) Get(manifestId string, periodId string, customXmlElementId string) (*model.CustomXmlElement, error) {
    var resp *model.CustomXmlElement
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["custom_xml_element_id"] = customXmlElementId
	}
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements/{custom_xml_element_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) Delete(manifestId string, periodId string, customXmlElementId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["custom_xml_element_id"] = customXmlElementId
	}
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements/{custom_xml_element_id}", &resp, reqParams)
    return resp, err
}
