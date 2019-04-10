package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMachineLearningObjectDetectionCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMachineLearningObjectDetectionCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMachineLearningObjectDetectionCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMachineLearningObjectDetectionCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMachineLearningObjectDetectionCustomdataApi) GetCustomData(encodingId string, objectDetectionId string) (*model.CustomData, error) {
    var resp *model.CustomData
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["object_detection_id"] = objectDetectionId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/customData", &resp, reqParams)
    return resp, err
}
