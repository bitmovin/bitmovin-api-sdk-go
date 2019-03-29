package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersEnhancedWatermarkCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersEnhancedWatermarkCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersEnhancedWatermarkCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersEnhancedWatermarkCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersEnhancedWatermarkCustomdataApi) GetCustomData(filterId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/enhanced-watermark/{filter_id}/customData", &resp, reqParams)
    return resp, err
}
