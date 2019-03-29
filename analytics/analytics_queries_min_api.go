package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesMinApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesMinApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesMinApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesMinApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesMinApi) Create(analyticsMinQueryRequest model.AnalyticsResponse) (*model.AnalyticsResponse, error) {
    payload := model.AnalyticsResponse(analyticsMinQueryRequest)
    
    err := api.apiClient.Post("/analytics/queries/min", &payload)
    return &payload, err
}
