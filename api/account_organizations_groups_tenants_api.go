package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsGroupsTenantsAPI communicates with '/account/organizations/{organization_id}/groups/{group_id}/tenants' endpoints
type AccountOrganizationsGroupsTenantsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAccountOrganizationsGroupsTenantsAPI constructor for AccountOrganizationsGroupsTenantsAPI that takes options as argument
func NewAccountOrganizationsGroupsTenantsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsGroupsTenantsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountOrganizationsGroupsTenantsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsGroupsTenantsAPIWithClient constructor for AccountOrganizationsGroupsTenantsAPI that takes an APIClient as argument
func NewAccountOrganizationsGroupsTenantsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsGroupsTenantsAPI {
	a := &AccountOrganizationsGroupsTenantsAPI{apiClient: apiClient}
	return a
}

// Create Add Tenant to Group
func (api *AccountOrganizationsGroupsTenantsAPI) Create(organizationId string, groupId string, tenant model.Tenant) (*model.Tenant, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
	}

	var responseModel model.Tenant
	err := api.apiClient.Post("/account/organizations/{organization_id}/groups/{group_id}/tenants", &tenant, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Tenant
func (api *AccountOrganizationsGroupsTenantsAPI) Delete(organizationId string, groupId string, tenantId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
		params.PathParams["tenant_id"] = tenantId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}/tenants/{tenant_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Tenant Details
func (api *AccountOrganizationsGroupsTenantsAPI) Get(organizationId string, groupId string, tenantId string) (*model.Tenant, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
		params.PathParams["tenant_id"] = tenantId
	}

	var responseModel model.Tenant
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/tenants/{tenant_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Tenants
func (api *AccountOrganizationsGroupsTenantsAPI) List(organizationId string, groupId string) (*pagination.TenantsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
	}

	var responseModel pagination.TenantsListPagination
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/tenants", nil, &responseModel, reqParams)
	return &responseModel, err
}
