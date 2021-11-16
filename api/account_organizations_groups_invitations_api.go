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

// Delete Invitation
func (api *AccountOrganizationsGroupsInvitationsAPI) Delete(organizationId string, groupId string, invitationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
		params.PathParams["invitation_id"] = invitationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}/invitations/{invitation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Invitation Details
func (api *AccountOrganizationsGroupsInvitationsAPI) Get(organizationId string, groupId string, invitationId string) (*model.Invitation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
		params.PathParams["invitation_id"] = invitationId
	}

	var responseModel model.Invitation
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/invitations/{invitation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Invitations
func (api *AccountOrganizationsGroupsInvitationsAPI) List(organizationId string, groupId string, queryParams ...func(*AccountOrganizationsGroupsInvitationsAPIListQueryParams)) (*pagination.InvitationsListPagination, error) {
	queryParameters := &AccountOrganizationsGroupsInvitationsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
		params.PathParams["group_id"] = groupId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.InvitationsListPagination
	err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/invitations", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AccountOrganizationsGroupsInvitationsAPIListQueryParams contains all query parameters for the List endpoint
type AccountOrganizationsGroupsInvitationsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AccountOrganizationsGroupsInvitationsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
