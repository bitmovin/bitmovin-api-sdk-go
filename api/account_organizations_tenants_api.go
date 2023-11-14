package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsTenantsAPI communicates with '/account/organizations/{organization_id}/tenants' endpoints
type AccountOrganizationsTenantsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAccountOrganizationsTenantsAPI constructor for AccountOrganizationsTenantsAPI that takes options as argument
func NewAccountOrganizationsTenantsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsTenantsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountOrganizationsTenantsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsTenantsAPIWithClient constructor for AccountOrganizationsTenantsAPI that takes an APIClient as argument
func NewAccountOrganizationsTenantsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsTenantsAPI {
	a := &AccountOrganizationsTenantsAPI{apiClient: apiClient}
	return a
}

// List all Tenants with their groups for a given organization
func (api *AccountOrganizationsTenantsAPI) List(organizationId string) (*pagination.TenantWithGroupssListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel pagination.TenantWithGroupssListPagination
	err := api.apiClient.Get("/account/organizations/{organization_id}/tenants", nil, &responseModel, reqParams)
	return &responseModel, err
}
