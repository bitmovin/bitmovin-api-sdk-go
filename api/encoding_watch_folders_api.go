package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingWatchFoldersAPI communicates with '/encoding/watch-folders' endpoints
type EncodingWatchFoldersAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingWatchFoldersAPI constructor for EncodingWatchFoldersAPI that takes options as argument
func NewEncodingWatchFoldersAPI(options ...apiclient.APIClientOption) (*EncodingWatchFoldersAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingWatchFoldersAPIWithClient(apiClient), nil
}

// NewEncodingWatchFoldersAPIWithClient constructor for EncodingWatchFoldersAPI that takes an APIClient as argument
func NewEncodingWatchFoldersAPIWithClient(apiClient *apiclient.APIClient) *EncodingWatchFoldersAPI {
	a := &EncodingWatchFoldersAPI{apiClient: apiClient}
	return a
}

// Create Watch Folder
func (api *EncodingWatchFoldersAPI) Create(watchFolder model.WatchFolder) (*model.WatchFolder, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.WatchFolder
	err := api.apiClient.Post("/encoding/watch-folders", &watchFolder, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Watch Folder
func (api *EncodingWatchFoldersAPI) Delete(watchFolderId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["watch_folder_id"] = watchFolderId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/watch-folders/{watch_folder_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Watch Folder details
func (api *EncodingWatchFoldersAPI) Get(watchFolderId string) (*model.WatchFolder, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["watch_folder_id"] = watchFolderId
	}

	var responseModel model.WatchFolder
	err := api.apiClient.Get("/encoding/watch-folders/{watch_folder_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Watch Folders
func (api *EncodingWatchFoldersAPI) List(queryParams ...func(*EncodingWatchFoldersAPIListQueryParams)) (*pagination.WatchFoldersListPagination, error) {
	queryParameters := &EncodingWatchFoldersAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.WatchFoldersListPagination
	err := api.apiClient.Get("/encoding/watch-folders", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Start Watch Folder
func (api *EncodingWatchFoldersAPI) Start(watchFolderId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["watch_folder_id"] = watchFolderId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/watch-folders/{watch_folder_id}/start", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Stop Watch Folder
func (api *EncodingWatchFoldersAPI) Stop(watchFolderId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["watch_folder_id"] = watchFolderId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Post("/encoding/watch-folders/{watch_folder_id}/stop", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingWatchFoldersAPIListQueryParams contains all query parameters for the List endpoint
type EncodingWatchFoldersAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingWatchFoldersAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
