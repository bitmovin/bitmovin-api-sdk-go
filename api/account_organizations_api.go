package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsAPI communicates with '/account/organizations' endpoints
type AccountOrganizationsAPI struct {
	apiClient *apiclient.APIClient

	// SubOrganizations communicates with '/account/organizations/{organization_id}/sub-organizations' endpoints
	SubOrganizations *AccountOrganizationsSubOrganizationsAPI
	// Groups communicates with '/account/organizations/{organization_id}/groups' endpoints
	Groups *AccountOrganizationsGroupsAPI
}

// NewAccountOrganizationsAPI constructor for AccountOrganizationsAPI that takes options as argument
func NewAccountOrganizationsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountOrganizationsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsAPIWithClient constructor for AccountOrganizationsAPI that takes an APIClient as argument
func NewAccountOrganizationsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsAPI {
	a := &AccountOrganizationsAPI{apiClient: apiClient}
	a.SubOrganizations = NewAccountOrganizationsSubOrganizationsAPIWithClient(apiClient)
	a.Groups = NewAccountOrganizationsGroupsAPIWithClient(apiClient)

	return a
}

// Create Add Organization
func (api *AccountOrganizationsAPI) Create(organization model.Organization) (*model.Organization, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Organization
	err := api.apiClient.Post("/account/organizations", &organization, &responseModel, reqParams)
	return &responseModel, err
}

// Get Organization Details
func (api *AccountOrganizationsAPI) Get(organizationId string) (*model.Organization, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel model.Organization
	err := api.apiClient.Get("/account/organizations/{organization_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Organizations
func (api *AccountOrganizationsAPI) List() (*pagination.OrganizationsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel pagination.OrganizationsListPagination
	err := api.apiClient.Get("/account/organizations", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update Organization
func (api *AccountOrganizationsAPI) Update(organizationId string, updateOrganizationRequest model.UpdateOrganizationRequest) (*model.Organization, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel model.Organization
	err := api.apiClient.Put("/account/organizations/{organization_id}", &updateOrganizationRequest, &responseModel, reqParams)
	return &responseModel, err
}
