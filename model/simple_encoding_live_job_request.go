package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// SimpleEncodingLiveJobRequest model
type SimpleEncodingLiveJobRequest struct {
	Input *SimpleEncodingLiveJobInput `json:"input,omitempty"`
	// Please take a look at the [documentation](https://bitmovin.com/docs/encoding/articles/simple-encoding-api-live#outputs) (required)
	Outputs []SimpleEncodingLiveJobOutput `json:"outputs,omitempty"`
	// The cloud region that will be used for the live encoding. This value has to be set.
	CloudRegion SimpleEncodingLiveCloudRegion `json:"cloudRegion,omitempty"`
	// This property will be used for naming the encoding.
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals model SimpleEncodingLiveJobRequest from a JSON structure
func (m *SimpleEncodingLiveJobRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		Input       *SimpleEncodingLiveJobInput   `json:"input"`
		Outputs     json.RawMessage               `json:"outputs"`
		CloudRegion SimpleEncodingLiveCloudRegion `json:"cloudRegion"`
		Name        *string                       `json:"name"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result SimpleEncodingLiveJobRequest

	result.Input = data.Input
	result.CloudRegion = data.CloudRegion
	result.Name = data.Name

	allOfOutputs, err := UnmarshalSimpleEncodingLiveJobOutputSlice(bytes.NewBuffer(data.Outputs), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Outputs = allOfOutputs

	*m = result

	return nil
}
