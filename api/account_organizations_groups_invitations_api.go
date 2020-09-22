package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsGroupsInvitationsAPI communicates with '/account/organizations/{organization_id}/groups/{group_id}/invitations' endpoints
type AccountOrganizationsGroupsInvitationsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAccountOrganizationsGroupsInvitationsAPI constructor for AccountOrganizationsGroupsInvitationsAPI that takes options as argument
func NewAccountOrganizationsGroupsInvitationsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsGroupsInvitationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountOrganizationsGroupsInvitationsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsGroupsInvitationsAPIWithClient constructor for AccountOrganizationsGroupsInvitationsAPI that takes an APIClient as argument
func NewAccountOrganizationsGroupsInvitationsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsGroupsInvitationsAPI {
	a := &AccountOrganizationsGroupsInvitationsAPI{apiClient: apiClient}
	return a
}

// Create Add Invitation to Group
func (api *AccountOrganizationsGroupsInvitationsAPI) Create(organizationId string, groupId string, invitation model.Invitation) (*model.Invitation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
	}

	var responseModel model.Invitation
	err := api.apiClient.Post("/account/organizations/{organization_id}/groups/{group_id}/invitations", &invitation, &responseModel, reqParams)
	return &responseModel, err
}

// List Invitations
func (api *AccountOrganizationsGroupsInvitationsAPI) List(organizationId string, groupId string) (*pagination.InvitationsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
	}

	var responseModel pagination.InvitationsListPagination
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/invitations", nil, &responseModel, reqParams)
	return &responseModel, err
}
