package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesCountApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesCountApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesCountApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesCountApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesCountApi) Create(analyticsCountQueryRequest model.AnalyticsResponse) (*model.AnalyticsResponse, error) {
    payload := model.AnalyticsResponse(analyticsCountQueryRequest)
    
    err := api.apiClient.Post("/analytics/queries/count", &payload)
    return &payload, err
}
