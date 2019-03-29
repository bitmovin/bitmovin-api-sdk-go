package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesMedianApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesMedianApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesMedianApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesMedianApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesMedianApi) Create(analyticsMedianQueryRequest model.AnalyticsResponse) (*model.AnalyticsResponse, error) {
    payload := model.AnalyticsResponse(analyticsMedianQueryRequest)
    
    err := api.apiClient.Post("/analytics/queries/median", &payload)
    return &payload, err
}
