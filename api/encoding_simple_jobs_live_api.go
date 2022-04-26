package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingSimpleJobsLiveAPI communicates with '/encoding/simple/jobs/live' endpoints
type EncodingSimpleJobsLiveAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingSimpleJobsLiveAPI constructor for EncodingSimpleJobsLiveAPI that takes options as argument
func NewEncodingSimpleJobsLiveAPI(options ...apiclient.APIClientOption) (*EncodingSimpleJobsLiveAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingSimpleJobsLiveAPIWithClient(apiClient), nil
}

// NewEncodingSimpleJobsLiveAPIWithClient constructor for EncodingSimpleJobsLiveAPI that takes an APIClient as argument
func NewEncodingSimpleJobsLiveAPIWithClient(apiClient *apiclient.APIClient) *EncodingSimpleJobsLiveAPI {
	a := &EncodingSimpleJobsLiveAPI{apiClient: apiClient}
	return a
}

// Create a Simple Encoding Live Job
func (api *EncodingSimpleJobsLiveAPI) Create(simpleEncodingLiveJobRequest model.SimpleEncodingLiveJobRequest) (*model.SimpleEncodingLiveJobResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.SimpleEncodingLiveJobResponse
	err := api.apiClient.Post("/encoding/simple/jobs/live", &simpleEncodingLiveJobRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Get Simple Encoding Live Job details
// Get the details of a Simple Live Encoding Job.  Check out our [Simple Encoding API Live Documentation](https://bitmovin.com/docs/encoding/articles/simple-encoding-api-live) for additional information about the Simple Encoding API Live.
func (api *EncodingSimpleJobsLiveAPI) Get(simpleEncodingJobId string) (*model.SimpleEncodingLiveJobResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["simple_encoding_job_id"] = simpleEncodingJobId
	}

	var responseModel model.SimpleEncodingLiveJobResponse
	err := api.apiClient.Get("/encoding/simple/jobs/live/{simple_encoding_job_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}
