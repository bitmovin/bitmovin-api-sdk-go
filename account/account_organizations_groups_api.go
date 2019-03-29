package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsGroupsApi struct {
    apiClient *common.ApiClient
    Tenants *AccountOrganizationsGroupsTenantsApi
    Permissions *AccountOrganizationsGroupsPermissionsApi
}

func NewAccountOrganizationsGroupsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsGroupsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsGroupsApi{apiClient: apiClient}

    tenantsApi, err := NewAccountOrganizationsGroupsTenantsApi(configs...)
    api.Tenants = tenantsApi
    permissionsApi, err := NewAccountOrganizationsGroupsPermissionsApi(configs...)
    api.Permissions = permissionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsGroupsApi) Delete(organizationId string, groupId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
	}
    err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsGroupsApi) Get(organizationId string, groupId string) (*model.Group, error) {
    var resp *model.Group
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
	}
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsGroupsApi) List(organizationId string) (*pagination.GroupsListPagination, error) {
    var resp *pagination.GroupsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
	}
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsGroupsApi) Create(organizationId string, group model.Group) (*model.Group, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }
    payload := model.Group(group)
    
    err := api.apiClient.Post("/account/organizations/{organization_id}/groups", &payload, reqParams)
    return &payload, err
}
