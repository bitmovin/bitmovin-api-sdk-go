package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type PlayerCustomBuildsWebDownloadApi struct {
    apiClient *common.ApiClient
}

func NewPlayerCustomBuildsWebDownloadApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsWebDownloadApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsWebDownloadApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerCustomBuildsWebDownloadApi) Get(customBuildId string) (*model.CustomPlayerBuildDownload, error) {
    var resp *model.CustomPlayerBuildDownload
    reqParams := func(params *common.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
	}
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}/download", &resp, reqParams)
    return resp, err
}
