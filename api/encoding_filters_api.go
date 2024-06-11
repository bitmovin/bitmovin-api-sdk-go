package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersAPI communicates with '/encoding/filters' endpoints
type EncodingFiltersAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/filters/{filter_id}/type' endpoints
	Type *EncodingFiltersTypeAPI
	// Conform communicates with '/encoding/filters/conform' endpoints
	Conform *EncodingFiltersConformAPI
	// Watermark communicates with '/encoding/filters/watermark' endpoints
	Watermark *EncodingFiltersWatermarkAPI
	// AudioVolume communicates with '/encoding/filters/audio-volume' endpoints
	AudioVolume *EncodingFiltersAudioVolumeAPI
	// AzureSpeechToCaptions communicates with '/encoding/filters/azure-speech-to-captions' endpoints
	AzureSpeechToCaptions *EncodingFiltersAzureSpeechToCaptionsAPI
	// EnhancedWatermark communicates with '/encoding/filters/enhanced-watermark' endpoints
	EnhancedWatermark *EncodingFiltersEnhancedWatermarkAPI
	// Crop communicates with '/encoding/filters/crop' endpoints
	Crop *EncodingFiltersCropAPI
	// Rotate communicates with '/encoding/filters/rotate' endpoints
	Rotate *EncodingFiltersRotateAPI
	// Deinterlace communicates with '/encoding/filters/deinterlace' endpoints
	Deinterlace *EncodingFiltersDeinterlaceAPI
	// EnhancedDeinterlace communicates with '/encoding/filters/enhanced-deinterlace' endpoints
	EnhancedDeinterlace *EncodingFiltersEnhancedDeinterlaceAPI
	// AudioMix communicates with '/encoding/filters/audio-mix' endpoints
	AudioMix *EncodingFiltersAudioMixAPI
	// DenoiseHqdn3d communicates with '/encoding/filters/denoise-hqdn3d' endpoints
	DenoiseHqdn3d *EncodingFiltersDenoiseHqdn3dAPI
	// EbuR128SinglePass communicates with '/encoding/filters/ebu-r128-single-pass' endpoints
	EbuR128SinglePass *EncodingFiltersEbuR128SinglePassAPI
	// Text communicates with '/encoding/filters/text' endpoints
	Text *EncodingFiltersTextAPI
	// Interlace communicates with '/encoding/filters/interlace' endpoints
	Interlace *EncodingFiltersInterlaceAPI
	// Unsharp communicates with '/encoding/filters/unsharp' endpoints
	Unsharp *EncodingFiltersUnsharpAPI
	// Scale communicates with '/encoding/filters/scale' endpoints
	Scale *EncodingFiltersScaleAPI
}

// NewEncodingFiltersAPI constructor for EncodingFiltersAPI that takes options as argument
func NewEncodingFiltersAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingFiltersAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAPIWithClient constructor for EncodingFiltersAPI that takes an APIClient as argument
func NewEncodingFiltersAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAPI {
	a := &EncodingFiltersAPI{apiClient: apiClient}
	a.Type = NewEncodingFiltersTypeAPIWithClient(apiClient)
	a.Conform = NewEncodingFiltersConformAPIWithClient(apiClient)
	a.Watermark = NewEncodingFiltersWatermarkAPIWithClient(apiClient)
	a.AudioVolume = NewEncodingFiltersAudioVolumeAPIWithClient(apiClient)
	a.AzureSpeechToCaptions = NewEncodingFiltersAzureSpeechToCaptionsAPIWithClient(apiClient)
	a.EnhancedWatermark = NewEncodingFiltersEnhancedWatermarkAPIWithClient(apiClient)
	a.Crop = NewEncodingFiltersCropAPIWithClient(apiClient)
	a.Rotate = NewEncodingFiltersRotateAPIWithClient(apiClient)
	a.Deinterlace = NewEncodingFiltersDeinterlaceAPIWithClient(apiClient)
	a.EnhancedDeinterlace = NewEncodingFiltersEnhancedDeinterlaceAPIWithClient(apiClient)
	a.AudioMix = NewEncodingFiltersAudioMixAPIWithClient(apiClient)
	a.DenoiseHqdn3d = NewEncodingFiltersDenoiseHqdn3dAPIWithClient(apiClient)
	a.EbuR128SinglePass = NewEncodingFiltersEbuR128SinglePassAPIWithClient(apiClient)
	a.Text = NewEncodingFiltersTextAPIWithClient(apiClient)
	a.Interlace = NewEncodingFiltersInterlaceAPIWithClient(apiClient)
	a.Unsharp = NewEncodingFiltersUnsharpAPIWithClient(apiClient)
	a.Scale = NewEncodingFiltersScaleAPIWithClient(apiClient)

	return a
}

// Get Filter Details
func (api *EncodingFiltersAPI) Get(filterId string) (*model.Filter, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.Filter
	err := api.apiClient.Get("/encoding/filters/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Filters
func (api *EncodingFiltersAPI) List(queryParams ...func(*EncodingFiltersAPIListQueryParams)) (*pagination.FiltersListPagination, error) {
	queryParameters := &EncodingFiltersAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.FiltersListPagination
	err := api.apiClient.Get("/encoding/filters", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingFiltersAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
	Sort   string `query:"sort"`
	Type_  string `query:"type"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
