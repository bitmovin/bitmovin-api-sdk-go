package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingConfigurationsVideoMjpegAPI communicates with '/encoding/configurations/video/mjpeg' endpoints
type EncodingConfigurationsVideoMjpegAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/configurations/video/mjpeg/{configuration_id}/customData' endpoints
	Customdata *EncodingConfigurationsVideoMjpegCustomdataAPI
}

// NewEncodingConfigurationsVideoMjpegAPI constructor for EncodingConfigurationsVideoMjpegAPI that takes options as argument
func NewEncodingConfigurationsVideoMjpegAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoMjpegAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoMjpegAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoMjpegAPIWithClient constructor for EncodingConfigurationsVideoMjpegAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoMjpegAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoMjpegAPI {
	a := &EncodingConfigurationsVideoMjpegAPI{apiClient: apiClient}
	a.Customdata = NewEncodingConfigurationsVideoMjpegCustomdataAPIWithClient(apiClient)

	return a
}

// Create MJPEG Codec Configuration
func (api *EncodingConfigurationsVideoMjpegAPI) Create(mjpegVideoConfiguration model.MjpegVideoConfiguration) (*model.MjpegVideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.MjpegVideoConfiguration
	err := api.apiClient.Post("/encoding/configurations/video/mjpeg", &mjpegVideoConfiguration, &responseModel, reqParams)
	return &responseModel, err
}

// Delete MJPEG Codec Configuration
func (api *EncodingConfigurationsVideoMjpegAPI) Delete(configurationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/configurations/video/mjpeg/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get MJPEG Codec Configuration Details
func (api *EncodingConfigurationsVideoMjpegAPI) Get(configurationId string) (*model.MjpegVideoConfiguration, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.MjpegVideoConfiguration
	err := api.apiClient.Get("/encoding/configurations/video/mjpeg/{configuration_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List MJPEG Configurations
func (api *EncodingConfigurationsVideoMjpegAPI) List(queryParams ...func(*EncodingConfigurationsVideoMjpegAPIListQueryParams)) (*pagination.MjpegVideoConfigurationsListPagination, error) {
	queryParameters := &EncodingConfigurationsVideoMjpegAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.MjpegVideoConfigurationsListPagination
	err := api.apiClient.Get("/encoding/configurations/video/mjpeg", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingConfigurationsVideoMjpegAPIListQueryParams contains all query parameters for the List endpoint
type EncodingConfigurationsVideoMjpegAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingConfigurationsVideoMjpegAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
