package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsGroupsAPI communicates with '/account/organizations/{organization_id}/groups' endpoints
type AccountOrganizationsGroupsAPI struct {
	apiClient *apiclient.APIClient

	// Tenants communicates with '/account/organizations/{organization_id}/groups/{group_id}/tenants' endpoints
	Tenants *AccountOrganizationsGroupsTenantsAPI
	// Invitations communicates with '/account/organizations/{organization_id}/groups/{group_id}/invitations' endpoints
	Invitations *AccountOrganizationsGroupsInvitationsAPI
	// Permissions communicates with '/account/organizations/{organization_id}/groups/{group_id}/permissions' endpoints
	Permissions *AccountOrganizationsGroupsPermissionsAPI
}

// NewAccountOrganizationsGroupsAPI constructor for AccountOrganizationsGroupsAPI that takes options as argument
func NewAccountOrganizationsGroupsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsGroupsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountOrganizationsGroupsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsGroupsAPIWithClient constructor for AccountOrganizationsGroupsAPI that takes an APIClient as argument
func NewAccountOrganizationsGroupsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsGroupsAPI {
	a := &AccountOrganizationsGroupsAPI{apiClient: apiClient}
	a.Tenants = NewAccountOrganizationsGroupsTenantsAPIWithClient(apiClient)
	a.Invitations = NewAccountOrganizationsGroupsInvitationsAPIWithClient(apiClient)
	a.Permissions = NewAccountOrganizationsGroupsPermissionsAPIWithClient(apiClient)

	return a
}

// Create Add Group
func (api *AccountOrganizationsGroupsAPI) Create(organizationId string, group model.Group) (*model.Group, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel model.Group
	err := api.apiClient.Post("/account/organizations/{organization_id}/groups", &group, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Group
func (api *AccountOrganizationsGroupsAPI) Delete(organizationId string, groupId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Group Details
func (api *AccountOrganizationsGroupsAPI) Get(organizationId string, groupId string) (*model.Group, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
	}

	var responseModel model.Group
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Groups
func (api *AccountOrganizationsGroupsAPI) List(organizationId string, queryParams ...func(*AccountOrganizationsGroupsAPIListQueryParams)) (*pagination.GroupsListPagination, error) {
	queryParameters := &AccountOrganizationsGroupsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.GroupsListPagination
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AccountOrganizationsGroupsAPIListQueryParams contains all query parameters for the List endpoint
type AccountOrganizationsGroupsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AccountOrganizationsGroupsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
