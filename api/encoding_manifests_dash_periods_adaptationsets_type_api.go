package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingManifestsDashPeriodsAdaptationsetsTypeAPI communicates with '/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/type' endpoints
type EncodingManifestsDashPeriodsAdaptationsetsTypeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsDashPeriodsAdaptationsetsTypeAPI constructor for EncodingManifestsDashPeriodsAdaptationsetsTypeAPI that takes options as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsTypeAPI(options ...apiclient.APIClientOption) (*EncodingManifestsDashPeriodsAdaptationsetsTypeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsDashPeriodsAdaptationsetsTypeAPIWithClient(apiClient), nil
}

// NewEncodingManifestsDashPeriodsAdaptationsetsTypeAPIWithClient constructor for EncodingManifestsDashPeriodsAdaptationsetsTypeAPI that takes an APIClient as argument
func NewEncodingManifestsDashPeriodsAdaptationsetsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsDashPeriodsAdaptationsetsTypeAPI {
	a := &EncodingManifestsDashPeriodsAdaptationsetsTypeAPI{apiClient: apiClient}
	return a
}

// Get adaptation set type
func (api *EncodingManifestsDashPeriodsAdaptationsetsTypeAPI) Get(manifestId string, periodId string, adaptationsetId string) (*model.AdaptationSetTypeResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["period_id"] = periodId
		params.PathParams["adaptationset_id"] = adaptationsetId
	}

	var responseModel model.AdaptationSetTypeResponse
	err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/type", nil, &responseModel, reqParams)
	return &responseModel, err
}
