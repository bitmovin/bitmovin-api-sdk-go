package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingOutputsAPI communicates with '/encoding/outputs' endpoints
type EncodingOutputsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/outputs/{output_id}/type' endpoints
	Type *EncodingOutputsTypeAPI
	// S3 communicates with '/encoding/outputs/s3' endpoints
	S3 *EncodingOutputsS3API
	// S3RoleBased communicates with '/encoding/outputs/s3-role-based' endpoints
	S3RoleBased *EncodingOutputsS3RoleBasedAPI
	// GenericS3 communicates with '/encoding/outputs/generic-s3' endpoints
	GenericS3 *EncodingOutputsGenericS3API
	// Local communicates with '/encoding/outputs/local' endpoints
	Local *EncodingOutputsLocalAPI
	// Gcs communicates with '/encoding/outputs/gcs' endpoints
	Gcs *EncodingOutputsGcsAPI
	// GcsServiceAccount communicates with '/encoding/outputs/gcs-service-account' endpoints
	GcsServiceAccount *EncodingOutputsGcsServiceAccountAPI
	// Azure communicates with '/encoding/outputs/azure' endpoints
	Azure *EncodingOutputsAzureAPI
	// Ftp communicates with '/encoding/outputs/ftp' endpoints
	Ftp *EncodingOutputsFtpAPI
	// Sftp communicates with '/encoding/outputs/sftp' endpoints
	Sftp *EncodingOutputsSftpAPI
	// AkamaiMsl communicates with '/encoding/outputs/akamai-msl' endpoints
	AkamaiMsl *EncodingOutputsAkamaiMslAPI
	// AkamaiNetstorage communicates with '/encoding/outputs/akamai-netstorage' endpoints
	AkamaiNetstorage *EncodingOutputsAkamaiNetstorageAPI
	// LiveMediaIngest communicates with '/encoding/outputs/live-media-ingest' endpoints
	LiveMediaIngest *EncodingOutputsLiveMediaIngestAPI
	// Cdn communicates with '/encoding/outputs/cdn' endpoints
	Cdn *EncodingOutputsCdnAPI
}

// NewEncodingOutputsAPI constructor for EncodingOutputsAPI that takes options as argument
func NewEncodingOutputsAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingOutputsAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAPIWithClient constructor for EncodingOutputsAPI that takes an APIClient as argument
func NewEncodingOutputsAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAPI {
	a := &EncodingOutputsAPI{apiClient: apiClient}
	a.Type = NewEncodingOutputsTypeAPIWithClient(apiClient)
	a.S3 = NewEncodingOutputsS3APIWithClient(apiClient)
	a.S3RoleBased = NewEncodingOutputsS3RoleBasedAPIWithClient(apiClient)
	a.GenericS3 = NewEncodingOutputsGenericS3APIWithClient(apiClient)
	a.Local = NewEncodingOutputsLocalAPIWithClient(apiClient)
	a.Gcs = NewEncodingOutputsGcsAPIWithClient(apiClient)
	a.GcsServiceAccount = NewEncodingOutputsGcsServiceAccountAPIWithClient(apiClient)
	a.Azure = NewEncodingOutputsAzureAPIWithClient(apiClient)
	a.Ftp = NewEncodingOutputsFtpAPIWithClient(apiClient)
	a.Sftp = NewEncodingOutputsSftpAPIWithClient(apiClient)
	a.AkamaiMsl = NewEncodingOutputsAkamaiMslAPIWithClient(apiClient)
	a.AkamaiNetstorage = NewEncodingOutputsAkamaiNetstorageAPIWithClient(apiClient)
	a.LiveMediaIngest = NewEncodingOutputsLiveMediaIngestAPIWithClient(apiClient)
	a.Cdn = NewEncodingOutputsCdnAPIWithClient(apiClient)

	return a
}

// Get Output Details
func (api *EncodingOutputsAPI) Get(outputId string) (*model.Output, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.Output
	err := api.apiClient.Get("/encoding/outputs/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Outputs
func (api *EncodingOutputsAPI) List(queryParams ...func(*EncodingOutputsAPIListQueryParams)) (*pagination.OutputsListPagination, error) {
	queryParameters := &EncodingOutputsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.OutputsListPagination
	err := api.apiClient.Get("/encoding/outputs", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingOutputsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingOutputsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingOutputsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
