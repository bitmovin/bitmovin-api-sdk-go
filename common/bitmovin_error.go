package common

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"encoding/json"
	"fmt"
)

type BitmovinError struct {
	Code             *int32
	Message          *string
	DeveloperMessage *string
	Details          []model.Message
	Links            []model.Link
}

func (e BitmovinError) Error() string {
	pretty, _ := json.MarshalIndent(e, "", "    ")
	return fmt.Sprintf("%s", string(pretty))
}
