package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsVirtualLicensesAPI communicates with '/analytics/virtual-licenses' endpoints
type AnalyticsVirtualLicensesAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsVirtualLicensesAPI constructor for AnalyticsVirtualLicensesAPI that takes options as argument
func NewAnalyticsVirtualLicensesAPI(options ...apiclient.APIClientOption) (*AnalyticsVirtualLicensesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsVirtualLicensesAPIWithClient(apiClient), nil
}

// NewAnalyticsVirtualLicensesAPIWithClient constructor for AnalyticsVirtualLicensesAPI that takes an APIClient as argument
func NewAnalyticsVirtualLicensesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsVirtualLicensesAPI {
	a := &AnalyticsVirtualLicensesAPI{apiClient: apiClient}
	return a
}

// Create Virtual License
func (api *AnalyticsVirtualLicensesAPI) Create(virtualLicenseCreateRequest model.VirtualLicenseCreateRequest) (*model.VirtualLicense, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.VirtualLicense
	err := api.apiClient.Post("/analytics/virtual-licenses", &virtualLicenseCreateRequest, &responseModel, reqParams)
	return &responseModel, err
}
