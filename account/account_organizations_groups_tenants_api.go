package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsGroupsTenantsApi struct {
    apiClient *common.ApiClient
}

func NewAccountOrganizationsGroupsTenantsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsGroupsTenantsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsGroupsTenantsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsGroupsTenantsApi) Delete(organizationId string, groupId string, tenantId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
        params.PathParams["tenant_id"] = tenantId
	}
    err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}/tenants/{tenant_id}", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsGroupsTenantsApi) Create(organizationId string, groupId string, tenant model.Tenant) (*model.Tenant, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }
    payload := model.Tenant(tenant)
    
    err := api.apiClient.Post("/account/organizations/{organization_id}/groups/{group_id}/tenants", &payload, reqParams)
    return &payload, err
}
func (api *AccountOrganizationsGroupsTenantsApi) Get(organizationId string, groupId string, tenantId string) (*model.Tenant, error) {
    var resp *model.Tenant
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
        params.PathParams["tenant_id"] = tenantId
	}
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/tenants/{tenant_id}", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsGroupsTenantsApi) List(organizationId string, groupId string) (*pagination.TenantsListPagination, error) {
    var resp *pagination.TenantsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
	}
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/tenants", &resp, reqParams)
    return resp, err
}
