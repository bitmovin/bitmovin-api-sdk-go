package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI struct {
	apiClient *apiclient.APIClient

	// Contentprotection communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}/contentprotection' endpoints
	Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPI
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI{apiClient: apiClient}
	a.Contentprotection = NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionAPIWithClient(apiClient)

	return a
}

// Create Add CMAF Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI) Create(manifestId string, periodId string, adaptationsetId string, dashCmafRepresentation model.DashCmafRepresentation) (*model.DashCmafRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.DashCmafRepresentation
	err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf", &dashCmafRepresentation, &responseModel, reqParams)
	return &responseModel, err
}

// Delete CMAF Representation
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get CMAF Representation Details
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashCmafRepresentation, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
	}

	var responseModel model.DashCmafRepresentation
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all CMAF Representations
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPI) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIListQueryParams)) (*pagination.DashCmafRepresentationsListPagination, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DashCmafRepresentationsListPagination
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
