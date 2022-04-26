package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// SimpleEncodingLiveJobResponse model
type SimpleEncodingLiveJobResponse struct {
	// The identifier of the Simple Encoding Live Job
	Id *string `json:"id,omitempty"`
	// The current status of the Simple Encoding Live Job
	Status SimpleEncodingLiveJobStatus `json:"status,omitempty"`
	// The identifier of the encoding that has been created based on the job request. This is only returned once the job execution has been successful and the encoding could be started.
	EncodingId *string `json:"encodingId,omitempty"`
	// The IP address of the encoder for this job request. This is only returned once the job execution has been successful and the encoding could be started. Ingest is expected to be sent to this IP address.
	EncoderIp *string `json:"encoderIp,omitempty"`
	// Stream key of the live encoder
	StreamKey *string                       `json:"streamKey,omitempty"`
	Input     *SimpleEncodingLiveJobInput   `json:"input,omitempty"`
	Outputs   []SimpleEncodingLiveJobOutput `json:"outputs,omitempty"`
	// Describes all the errors in cases the status of the job is 'error'.
	Errors []SimpleEncodingVodJobErrors `json:"errors,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ. The job is updated for example when the status changes
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// This property will be used for naming the encoding and the manifests.
	Name *string `json:"name,omitempty"`
	// The cloud region that will be used for the live encoding
	CloudRegion SimpleEncodingLiveCloudRegion `json:"cloudRegion,omitempty"`
}

// UnmarshalJSON unmarshals model SimpleEncodingLiveJobResponse from a JSON structure
func (m *SimpleEncodingLiveJobResponse) UnmarshalJSON(raw []byte) error {
	var data struct {
		Id          *string                       `json:"id"`
		Status      SimpleEncodingLiveJobStatus   `json:"status"`
		EncodingId  *string                       `json:"encodingId"`
		EncoderIp   *string                       `json:"encoderIp"`
		StreamKey   *string                       `json:"streamKey"`
		Input       *SimpleEncodingLiveJobInput   `json:"input"`
		Outputs     json.RawMessage               `json:"outputs"`
		Errors      []SimpleEncodingVodJobErrors  `json:"errors"`
		CreatedAt   *DateTime                     `json:"createdAt"`
		ModifiedAt  *DateTime                     `json:"modifiedAt"`
		Name        *string                       `json:"name"`
		CloudRegion SimpleEncodingLiveCloudRegion `json:"cloudRegion"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result SimpleEncodingLiveJobResponse

	result.Id = data.Id
	result.Status = data.Status
	result.EncodingId = data.EncodingId
	result.EncoderIp = data.EncoderIp
	result.StreamKey = data.StreamKey
	result.Input = data.Input
	result.Errors = data.Errors
	result.CreatedAt = data.CreatedAt
	result.ModifiedAt = data.ModifiedAt
	result.Name = data.Name
	result.CloudRegion = data.CloudRegion

	allOfOutputs, err := UnmarshalSimpleEncodingLiveJobOutputSlice(bytes.NewBuffer(data.Outputs), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Outputs = allOfOutputs

	*m = result

	return nil
}
