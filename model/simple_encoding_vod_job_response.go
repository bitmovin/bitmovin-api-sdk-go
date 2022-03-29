package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// SimpleEncodingVodJobResponse model
type SimpleEncodingVodJobResponse struct {
	// The identifier of the Simple Encoding VOD Job
	Id *string `json:"id,omitempty"`
	// The current status of the Simple Encoding VOD Job
	Status SimpleEncodingVodJobStatus `json:"status,omitempty"`
	// The template that has been used for the encoding.
	EncodingTemplate EncodingTemplate `json:"encodingTemplate,omitempty"`
	// The identifier of the encoding that has been created based on the job request. This is only returned once the job execution has been successful and the encoding could be started.
	EncodingId *string                        `json:"encodingId,omitempty"`
	Inputs     []SimpleEncodingVodJobUrlInput `json:"inputs,omitempty"`
	Outputs    []SimpleEncodingVodJobOutput   `json:"outputs,omitempty"`
	// Describes all the errors in cases the status of the job is 'error'.
	Errors []SimpleEncodingVodJobErrors `json:"errors,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ. The job is updated for example when the status changes
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals model SimpleEncodingVodJobResponse from a JSON structure
func (m *SimpleEncodingVodJobResponse) UnmarshalJSON(raw []byte) error {
	var data struct {
		Id               *string                        `json:"id"`
		Status           SimpleEncodingVodJobStatus     `json:"status"`
		EncodingTemplate EncodingTemplate               `json:"encodingTemplate"`
		EncodingId       *string                        `json:"encodingId"`
		Inputs           []SimpleEncodingVodJobUrlInput `json:"inputs"`
		Outputs          json.RawMessage                `json:"outputs"`
		Errors           []SimpleEncodingVodJobErrors   `json:"errors"`
		CreatedAt        *DateTime                      `json:"createdAt"`
		ModifiedAt       *DateTime                      `json:"modifiedAt"`
		Name             *string                        `json:"name"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result SimpleEncodingVodJobResponse

	result.Id = data.Id
	result.Status = data.Status
	result.EncodingTemplate = data.EncodingTemplate
	result.EncodingId = data.EncodingId
	result.Inputs = data.Inputs
	result.Errors = data.Errors
	result.CreatedAt = data.CreatedAt
	result.ModifiedAt = data.ModifiedAt
	result.Name = data.Name

	allOfOutputs, err := UnmarshalSimpleEncodingVodJobOutputSlice(bytes.NewBuffer(data.Outputs), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Outputs = allOfOutputs

	*m = result

	return nil
}
