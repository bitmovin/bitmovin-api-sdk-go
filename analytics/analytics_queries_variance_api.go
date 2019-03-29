package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesVarianceApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesVarianceApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesVarianceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesVarianceApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesVarianceApi) Create(analyticsVarianceQueryRequest model.AnalyticsResponse) (*model.AnalyticsResponse, error) {
    payload := model.AnalyticsResponse(analyticsVarianceQueryRequest)
    
    err := api.apiClient.Post("/analytics/queries/variance", &payload)
    return &payload, err
}
