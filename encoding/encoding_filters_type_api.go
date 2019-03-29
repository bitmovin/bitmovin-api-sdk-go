package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersTypeApi(configs ...func(*common.ApiClient)) (*EncodingFiltersTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersTypeApi) Get(filterId string) (*model.FilterTypeResponse, error) {
    var resp *model.FilterTypeResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/{filter_id}/type", &resp, reqParams)
    return resp, err
}
