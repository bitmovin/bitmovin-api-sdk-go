package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsLiveEsamMediaPointsAPI communicates with '/encoding/encodings/{encoding_id}/live/esam/media-points' endpoints
type EncodingEncodingsLiveEsamMediaPointsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsLiveEsamMediaPointsAPI constructor for EncodingEncodingsLiveEsamMediaPointsAPI that takes options as argument
func NewEncodingEncodingsLiveEsamMediaPointsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsLiveEsamMediaPointsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsLiveEsamMediaPointsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsLiveEsamMediaPointsAPIWithClient constructor for EncodingEncodingsLiveEsamMediaPointsAPI that takes an APIClient as argument
func NewEncodingEncodingsLiveEsamMediaPointsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsLiveEsamMediaPointsAPI {
	a := &EncodingEncodingsLiveEsamMediaPointsAPI{apiClient: apiClient}
	return a
}

// Create ESAM media point for live stream
// Creates an out-of-band Event Signaling and Management (ESAM) media point for insertion into a  running live stream at a specific match time. The media point contains signals with timing  offsets and optional conditions for content boundary signaling. To use this endpoint, ESAM media point insertion needs to be enabled in the &#x60;adInsertionSettings&#x60;  of the [StartLiveEncodingRequest](#/Encoding/PostEncodingEncodingsLiveStartByEncodingId).  **Out-of-band vs In-band ESAM Processing:**  - **Out-of-band (this endpoint):** Manually trigger SCTE-35 signals by creating media points    during a live stream. Use this when you want programmatic control over when ad breaks or    content boundaries occur, independent of the incoming stream content.  - **In-band (&#x60;EsamSettings&#x60;):** Automatically process SCTE-35 markers already present in the    incoming live stream. Configure EsamSettings in the    [StartLiveEncodingRequest](#/Encoding/PostEncodingEncodingsLiveStartByEncodingId) to connect    to a POIS endpoint that processes detected SCTE-35 messages.  Use out-of-band media points when you need to insert signals at specific times based on external triggers or business logic, rather than relying on SCTE-35 markers present in the source stream.
func (api *EncodingEncodingsLiveEsamMediaPointsAPI) Create(encodingId string, esamMediaPoint model.EsamMediaPoint) (*model.EsamMediaPoint, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.EsamMediaPoint
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/esam/media-points", &esamMediaPoint, &responseModel, reqParams)
	return &responseModel, err
}
