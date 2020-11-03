package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountOrganizationsGroupsPermissionsAPI communicates with '/account/organizations/{organization_id}/groups/{group_id}/permissions' endpoints
type AccountOrganizationsGroupsPermissionsAPI struct {
    apiClient *apiclient.APIClient
}

// NewAccountOrganizationsGroupsPermissionsAPI constructor for AccountOrganizationsGroupsPermissionsAPI that takes options as argument
func NewAccountOrganizationsGroupsPermissionsAPI(options ...apiclient.APIClientOption) (*AccountOrganizationsGroupsPermissionsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAccountOrganizationsGroupsPermissionsAPIWithClient(apiClient), nil
}

// NewAccountOrganizationsGroupsPermissionsAPIWithClient constructor for AccountOrganizationsGroupsPermissionsAPI that takes an APIClient as argument
func NewAccountOrganizationsGroupsPermissionsAPIWithClient(apiClient *apiclient.APIClient) *AccountOrganizationsGroupsPermissionsAPI {
    a := &AccountOrganizationsGroupsPermissionsAPI{apiClient: apiClient}
    return a
}

// Create Set Group Permissions
func (api *AccountOrganizationsGroupsPermissionsAPI) Create(organizationId string, groupId string, acl model.Acl) (*model.Acl, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }

    var responseModel model.Acl
    err := api.apiClient.Post("/account/organizations/{organization_id}/groups/{group_id}/permissions", &acl, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Permission
func (api *AccountOrganizationsGroupsPermissionsAPI) Delete(organizationId string, groupId string, permissionId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
        params.PathParams["permission_id"] = permissionId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}/permissions/{permission_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Get Group Permissions
func (api *AccountOrganizationsGroupsPermissionsAPI) List(organizationId string, groupId string) (*pagination.AclsListPagination, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }

    var responseModel pagination.AclsListPagination
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/permissions", nil, &responseModel, reqParams)
    return &responseModel, err
}

