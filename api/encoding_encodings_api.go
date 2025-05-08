package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsAPI communicates with '/encoding/encodings' endpoints
type EncodingEncodingsAPI struct {
	apiClient *apiclient.APIClient

	// Live communicates with '/encoding/encodings/{encoding_id}/live' endpoints
	Live *EncodingEncodingsLiveAPI
	// Customdata communicates with '/encoding/encodings/{encoding_id}/customData' endpoints
	Customdata *EncodingEncodingsCustomdataAPI
	// Streams communicates with '/encoding/encodings/{encoding_id}/streams' endpoints
	Streams *EncodingEncodingsStreamsAPI
	// InputStreams communicates with '/encoding/encodings/{encoding_id}/input-streams' endpoints
	InputStreams *EncodingEncodingsInputStreamsAPI
	// Muxings communicates with '/encoding/encodings/{encoding_id}/muxings' endpoints
	Muxings *EncodingEncodingsMuxingsAPI
	// Template communicates with '/encoding/encodings/{encoding_id}/template' endpoints
	Template *EncodingEncodingsTemplateAPI
	// TransferRetries communicates with '/encoding/encodings/{encoding_id}/transfer-retries' endpoints
	TransferRetries *EncodingEncodingsTransferRetriesAPI
	// OutputPaths communicates with '/encoding/encodings/{encoding_id}/output-paths' endpoints
	OutputPaths *EncodingEncodingsOutputPathsAPI
	// Captions intermediary API object with no endpoints
	Captions *EncodingEncodingsCaptionsAPI
	// Sidecars communicates with '/encoding/encodings/{encoding_id}/sidecars' endpoints
	Sidecars *EncodingEncodingsSidecarsAPI
	// Keyframes communicates with '/encoding/encodings/{encoding_id}/keyframes' endpoints
	Keyframes *EncodingEncodingsKeyframesAPI
	// Scte35Triggers communicates with '/encoding/encodings/{encoding_id}/scte-35-triggers' endpoints
	Scte35Triggers *EncodingEncodingsScte35TriggersAPI
}

// NewEncodingEncodingsAPI constructor for EncodingEncodingsAPI that takes options as argument
func NewEncodingEncodingsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsAPIWithClient constructor for EncodingEncodingsAPI that takes an APIClient as argument
func NewEncodingEncodingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsAPI {
	a := &EncodingEncodingsAPI{apiClient: apiClient}
	a.Live = NewEncodingEncodingsLiveAPIWithClient(apiClient)
	a.Customdata = NewEncodingEncodingsCustomdataAPIWithClient(apiClient)
	a.Streams = NewEncodingEncodingsStreamsAPIWithClient(apiClient)
	a.InputStreams = NewEncodingEncodingsInputStreamsAPIWithClient(apiClient)
	a.Muxings = NewEncodingEncodingsMuxingsAPIWithClient(apiClient)
	a.Template = NewEncodingEncodingsTemplateAPIWithClient(apiClient)
	a.TransferRetries = NewEncodingEncodingsTransferRetriesAPIWithClient(apiClient)
	a.OutputPaths = NewEncodingEncodingsOutputPathsAPIWithClient(apiClient)
	a.Captions = NewEncodingEncodingsCaptionsAPIWithClient(apiClient)
	a.Sidecars = NewEncodingEncodingsSidecarsAPIWithClient(apiClient)
	a.Keyframes = NewEncodingEncodingsKeyframesAPIWithClient(apiClient)
	a.Scte35Triggers = NewEncodingEncodingsScte35TriggersAPIWithClient(apiClient)

	return a
}

// Create Encoding
func (api *EncodingEncodingsAPI) Create(encoding model.Encoding) (*model.Encoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Encoding
	err := api.apiClient.Post("/encoding/encodings", &encoding, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Encoding
func (api *EncodingEncodingsAPI) Delete(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Encoding Details
func (api *EncodingEncodingsAPI) Get(encodingId string) (*model.Encoding, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Encoding
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// GetStartRequest Encoding Start Details
func (api *EncodingEncodingsAPI) GetStartRequest(encodingId string) (*model.StartEncodingRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.StartEncodingRequest
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Encodings
// Bitmovin retains historical encoding jobs data for a period of 90 days. Should you require the data to be stored for an extended duration, it is necessary to store it within your own data repository.
func (api *EncodingEncodingsAPI) List(queryParams ...func(*EncodingEncodingsAPIListQueryParams)) (*pagination.EncodingsListPagination, error) {
	queryParameters := &EncodingEncodingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EncodingsListPagination
	err := api.apiClient.Get("/encoding/encodings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Reprioritize Encoding
func (api *EncodingEncodingsAPI) Reprioritize(encodingId string, reprioritizeEncodingRequest model.ReprioritizeEncodingRequest) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/reprioritize", &reprioritizeEncodingRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Reschedule Encoding
func (api *EncodingEncodingsAPI) Reschedule(encodingId string, rescheduleEncodingRequest model.RescheduleEncodingRequest) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/reschedule", &rescheduleEncodingRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Start VoD Encoding
func (api *EncodingEncodingsAPI) Start(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start VoD Encoding
func (api *EncodingEncodingsAPI) StartWithRequestBody(encodingId string, startEncodingRequest model.StartEncodingRequest) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/start", &startEncodingRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Status Encoding Status
func (api *EncodingEncodingsAPI) Status(encodingId string) (*model.ModelTask, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ModelTask
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/status", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop Encoding
func (api *EncodingEncodingsAPI) Stop(encodingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsAPIListQueryParams struct {
	Offset                 int32              `query:"offset"`
	Limit                  int32              `query:"limit"`
	IncludeTotalCount      bool               `query:"includeTotalCount"`
	Sort                   string             `query:"sort"`
	Type_                  string             `query:"type"`
	Status                 string             `query:"status"`
	CloudRegion            model.CloudRegion  `query:"cloudRegion"`
	SelectedCloudRegion    model.CloudRegion  `query:"selectedCloudRegion"`
	EncoderVersion         string             `query:"encoderVersion"`
	SelectedEncoderVersion string             `query:"selectedEncoderVersion"`
	SelectedEncodingMode   model.EncodingMode `query:"selectedEncodingMode"`
	Name                   string             `query:"name"`
	Search                 string             `query:"search"`
	CreatedAtNewerThan     model.DateTime     `query:"createdAtNewerThan"`
	CreatedAtOlderThan     model.DateTime     `query:"createdAtOlderThan"`
	StartedAtNewerThan     model.DateTime     `query:"startedAtNewerThan"`
	StartedAtOlderThan     model.DateTime     `query:"startedAtOlderThan"`
	FinishedAtNewerThan    model.DateTime     `query:"finishedAtNewerThan"`
	FinishedAtOlderThan    model.DateTime     `query:"finishedAtOlderThan"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
