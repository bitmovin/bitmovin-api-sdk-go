package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// SimpleEncodingVodJobRequest model
type SimpleEncodingVodJobRequest struct {
	// The template that will be used for the encoding.
	EncodingTemplate EncodingTemplate            `json:"encodingTemplate,omitempty"`
	Inputs           []SimpleEncodingVodJobInput `json:"inputs,omitempty"`
	// Please take a look at the [documentation](https://bitmovin.com/docs/encoding/articles/simple-encoding-api#inputs-outputs) (required)
	Outputs []SimpleEncodingVodJobOutput `json:"outputs,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals model SimpleEncodingVodJobRequest from a JSON structure
func (m *SimpleEncodingVodJobRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		EncodingTemplate EncodingTemplate `json:"encodingTemplate"`
		Inputs           json.RawMessage  `json:"inputs"`
		Outputs          json.RawMessage  `json:"outputs"`
		Name             *string          `json:"name"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result SimpleEncodingVodJobRequest

	result.EncodingTemplate = data.EncodingTemplate
	result.Name = data.Name

	allOfInputs, err := UnmarshalSimpleEncodingVodJobInputSlice(bytes.NewBuffer(data.Inputs), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Inputs = allOfInputs
	allOfOutputs, err := UnmarshalSimpleEncodingVodJobOutputSlice(bytes.NewBuffer(data.Outputs), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Outputs = allOfOutputs

	*m = result

	return nil
}
