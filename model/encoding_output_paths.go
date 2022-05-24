package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// EncodingOutputPaths model
type EncodingOutputPaths struct {
	Output *Output                       `json:"output,omitempty"`
	Paths  *EncodingOutputPathsForOutput `json:"paths,omitempty"`
}

// UnmarshalJSON unmarshals model EncodingOutputPaths from a JSON structure
func (m *EncodingOutputPaths) UnmarshalJSON(raw []byte) error {
	var data struct {
		Output json.RawMessage               `json:"output"`
		Paths  *EncodingOutputPathsForOutput `json:"paths"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result EncodingOutputPaths

	result.Paths = data.Paths

	allOfOutput, err := UnmarshalOutput(bytes.NewBuffer(data.Output), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Output = &allOfOutput

	*m = result

	return nil
}
