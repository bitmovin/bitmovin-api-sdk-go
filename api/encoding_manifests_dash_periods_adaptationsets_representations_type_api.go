package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/{representation_id}/type' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI{apiClient: apiClient}
	return a
}

// Get DASH representation type
func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPI) Get(manifestId string, periodId string, adaptationsetId string, representationId string, queryParams ...func(*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIGetQueryParams)) (*model.DashRepresentationTypeResponse, error) {
	queryParameters := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIGetQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
		params.PathParams["representation_id"] = representationId
		params.QueryParams = queryParameters
	}

	var responseModel model.DashRepresentationTypeResponse
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/{representation_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIGetQueryParams contains all query parameters for the Get endpoint
type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIGetQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsTypeAPIGetQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
