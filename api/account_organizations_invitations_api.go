package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsInvitationsAPI communicates with '/account/organizations/{organization_id}/invitations' endpoints
type AccountOrganizationsInvitationsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAccountOrganizationsInvitationsAPI constructor for AccountOrganizationsInvitationsAPI that takes options as argument
func NewAccountOrganizationsInvitationsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsInvitationsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountOrganizationsInvitationsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsInvitationsAPIWithClient constructor for AccountOrganizationsInvitationsAPI that takes an APIClient as argument
func NewAccountOrganizationsInvitationsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsInvitationsAPI {
	a := &AccountOrganizationsInvitationsAPI{apiClient: apiClient}
	return a
}

// List all pending invitations of an org id
func (api *AccountOrganizationsInvitationsAPI) List(organizationId string) (*pagination.OrganizationPendingInvitationsListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel pagination.OrganizationPendingInvitationsListPagination
	err := api.apiClient.Get("/account/organizations/{organization_id}/invitations", nil, &responseModel, reqParams)
	return &responseModel, err
}
