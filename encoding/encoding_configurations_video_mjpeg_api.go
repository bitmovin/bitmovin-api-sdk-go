package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoMjpegApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoMjpegCustomdataApi
}

func NewEncodingConfigurationsVideoMjpegApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoMjpegApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoMjpegApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoMjpegCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoMjpegApi) Get(configurationId string) (*model.MjpegVideoConfiguration, error) {
    var resp *model.MjpegVideoConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/video/mjpeg/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoMjpegApi) Create(mjpegVideoConfiguration model.MjpegVideoConfiguration) (*model.MjpegVideoConfiguration, error) {
    payload := model.MjpegVideoConfiguration(mjpegVideoConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/video/mjpeg", &payload)
    return &payload, err
}
func (api *EncodingConfigurationsVideoMjpegApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/video/mjpeg/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoMjpegApi) List(queryParams ...func(*query.MjpegVideoConfigurationListQueryParams)) (*pagination.MjpegVideoConfigurationsListPagination, error) {
    queryParameters := &query.MjpegVideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.MjpegVideoConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/video/mjpeg", &resp, reqParams)
    return resp, err
}
