package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersRotateCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersRotateCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersRotateCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersRotateCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersRotateCustomdataApi) GetCustomData(filterId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/rotate/{filter_id}/customData", &resp, reqParams)
    return resp, err
}
