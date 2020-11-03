package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsSubOrganizationsAPI communicates with '/account/organizations/{organization_id}/sub-organizations' endpoints
type AccountOrganizationsSubOrganizationsAPI struct {
    apiClient *apiclient.APIClient
}

// NewAccountOrganizationsSubOrganizationsAPI constructor for AccountOrganizationsSubOrganizationsAPI that takes options as argument
func NewAccountOrganizationsSubOrganizationsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsSubOrganizationsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAccountOrganizationsSubOrganizationsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsSubOrganizationsAPIWithClient constructor for AccountOrganizationsSubOrganizationsAPI that takes an APIClient as argument
func NewAccountOrganizationsSubOrganizationsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsSubOrganizationsAPI {
    a := &AccountOrganizationsSubOrganizationsAPI{apiClient: apiClient}
    return a
}

// List Organizations under given parent organization
func (api *AccountOrganizationsSubOrganizationsAPI) List(organizationId string) (*pagination.OrganizationsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel pagination.OrganizationsListPagination
    err := api.apiClient.Get("/account/organizations/{organization_id}/sub-organizations", nil, &responseModel, reqParams)
    return &responseModel, err
}

