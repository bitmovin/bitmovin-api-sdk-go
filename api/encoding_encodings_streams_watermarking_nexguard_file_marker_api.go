package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker' endpoints
type EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker/{nexguard_id}/customData' endpoints
	Customdata *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPI
}

// NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI constructor for EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI that takes options as argument
func NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIWithClient constructor for EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI {
	a := &EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsStreamsWatermarkingNexguardFileMarkerCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add a nexguard file marker watermarking configurations
// Nexguard FileMarker watermarking has several restrictions on the shape of your streams and muxings. The supported muxings are currently fMP4, TS and WebM; segment naming must follow &#x60;&lt;filename&gt;_&lt;number&gt;.&lt;extension&gt;&#x60;; init segment naming must follow &#x60;&lt;filename&gt;_init.&lt;extension&gt;&#x60; Supported framerates:   * 23.976   * 24.000   * 25.000   * 29.970   * 30.000   * 48.000   * 50.000   * 59.940   * 60.000  Resolution:   * 320 &lt;&#x3D; width &lt;&#x3D; 5120   * 240 &lt;&#x3D; height &lt;&#x3D; 3200  And the GOP size has to be 2 or 2.002 seconds. Please note that our api requires the gop size to be in frames.
func (api *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI) Create(encodingId string, streamId string, nexGuardFileMarker model.NexGuardFileMarker) (*model.NexGuardFileMarker, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.NexGuardFileMarker
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker", &nexGuardFileMarker, &responseModel, reqParams)
	return &responseModel, err
}

// Delete nexguard file marker watermarking configurations
func (api *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI) Delete(encodingId string, streamId string, nexguardId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["nexguard_id"] = nexguardId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker/{nexguard_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Nexguard file marker watermarking configurations details
func (api *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI) Get(encodingId string, streamId string, nexguardId string) (*model.NexGuardFileMarker, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["nexguard_id"] = nexguardId
	}

	var responseModel model.NexGuardFileMarker
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker/{nexguard_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List nexguard file marker watermarking configurations
func (api *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIListQueryParams)) (*pagination.NexGuardFileMarkersListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.NexGuardFileMarkersListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/watermarking/nexguard-file-marker", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsWatermarkingNexguardFileMarkerAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
