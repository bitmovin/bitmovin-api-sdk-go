package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsGroupsPermissionsApi struct {
    apiClient *common.ApiClient
}

func NewAccountOrganizationsGroupsPermissionsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsGroupsPermissionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsGroupsPermissionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsGroupsPermissionsApi) List(organizationId string, groupId string) (*pagination.AclsListPagination, error) {
    var resp *pagination.AclsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
	}
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/permissions", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsGroupsPermissionsApi) Create(organizationId string, groupId string, acl model.Acl) (*model.Acl, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }
    payload := model.Acl(acl)
    
    err := api.apiClient.Post("/account/organizations/{organization_id}/groups/{group_id}/permissions", &payload, reqParams)
    return &payload, err
}
func (api *AccountOrganizationsGroupsPermissionsApi) Delete(organizationId string, groupId string, permissionId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
        params.PathParams["permission_id"] = permissionId
	}
    err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}/permissions/{permission_id}", &resp, reqParams)
    return resp, err
}
