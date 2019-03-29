package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsApi struct {
    apiClient *common.ApiClient
    Groups *AccountOrganizationsGroupsApi
}

func NewAccountOrganizationsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsApi{apiClient: apiClient}

    groupsApi, err := NewAccountOrganizationsGroupsApi(configs...)
    api.Groups = groupsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsApi) List() (*pagination.OrganizationsListPagination, error) {
    var resp *pagination.OrganizationsListPagination
    reqParams := func(params *common.RequestParams) {
	}
    err := api.apiClient.Get("/account/organizations", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsApi) Get(organizationId string) (*model.Organization, error) {
    var resp *model.Organization
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
	}
    err := api.apiClient.Get("/account/organizations/{organization_id}", &resp, reqParams)
    return resp, err
}
func (api *AccountOrganizationsApi) Create(organization model.Organization) (*model.Organization, error) {
    payload := model.Organization(organization)
    
    err := api.apiClient.Post("/account/organizations", &payload)
    return &payload, err
}
func (api *AccountOrganizationsApi) Delete(organizationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
	}
    err := api.apiClient.Delete("/account/organizations/{organization_id}", &resp, reqParams)
    return resp, err
}
