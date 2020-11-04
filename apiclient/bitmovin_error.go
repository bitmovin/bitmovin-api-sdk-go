package apiclient

import (
	"encoding/json"
	"fmt"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type BitmovinError struct {
	ErrorCode        *int32          `json:"code,omitempty"`
	Message          string          `json:"message,omitempty"`
	ShortMessage     string          `json:"shortMessage,omitempty"`
	RequestId        string          `json:"requestId,omitempty"`
	DeveloperMessage string          `json:"developerMessage,omitempty"`
	Details          []model.Message `json:"details,omitempty"`
	Links            []model.Link    `json:"links,omitempty"`

	HttpStatusCode *int
	RawResponse    string
}

func (e BitmovinError) Error() string {
	pretty, _ := json.MarshalIndent(e, "", "    ")
	return fmt.Sprintf("%s", string(pretty))
}
