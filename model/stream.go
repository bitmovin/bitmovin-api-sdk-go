package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// Stream model
type Stream struct {
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Determines the input source(s) for the stream. All video streams of an encoding need to have identical input configurations (required)
	InputStreams []StreamInput    `json:"inputStreams,omitempty"`
	Outputs      []EncodingOutput `json:"outputs,omitempty"`
	// Set true to create quality metadata for this stream
	CreateQualityMetaData *bool `json:"createQualityMetaData,omitempty"`
	// Id of the codec configuration (required)
	CodecConfigId *string `json:"codecConfigId,omitempty"`
	// Number of encoded segments. Available after encoding finishes.
	SegmentsEncoded *int32 `json:"segmentsEncoded,omitempty"`
	// Conditions to evaluate before creating the stream. If this evaluation fails, the stream won't be created. All muxings that depend on the stream will also not be created.
	Conditions *AbstractCondition `json:"conditions,omitempty"`
	// If this is set and contains objects, then this stream has been ignored during the encoding process
	IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
	// Mode of the stream
	Mode StreamMode `json:"mode,omitempty"`
	// The encoding mode of the stream which was applied by the assigned codec configuration
	SelectedEncodingMode EncodingMode `json:"selectedEncodingMode,omitempty"`
	// Settings to configure Per-Title on stream level
	PerTitleSettings *StreamPerTitleSettings `json:"perTitleSettings,omitempty"`
	Metadata         *StreamMetadata         `json:"metadata,omitempty"`
	// Determines how to react to errors during decoding
	DecodingErrorMode DecodingErrorMode `json:"decodingErrorMode,omitempty"`
	// Contains stream properties which may not have been defined in the configuration
	AppliedSettings *AppliedStreamSettings `json:"appliedSettings,omitempty"`
}

// UnmarshalJSON unmarshals model Stream from a JSON structure
func (m *Stream) UnmarshalJSON(raw []byte) error {
	var data struct {
		Name                  *string                 `json:"name"`
		Description           *string                 `json:"description"`
		CreatedAt             *DateTime               `json:"createdAt"`
		ModifiedAt            *DateTime               `json:"modifiedAt"`
		CustomData            *map[string]interface{} `json:"customData"`
		Id                    *string                 `json:"id"`
		InputStreams          []StreamInput           `json:"inputStreams"`
		Outputs               []EncodingOutput        `json:"outputs"`
		CreateQualityMetaData *bool                   `json:"createQualityMetaData"`
		CodecConfigId         *string                 `json:"codecConfigId"`
		SegmentsEncoded       *int32                  `json:"segmentsEncoded"`
		Conditions            json.RawMessage         `json:"conditions"`
		IgnoredBy             []Ignoring              `json:"ignoredBy"`
		Mode                  StreamMode              `json:"mode"`
		SelectedEncodingMode  EncodingMode            `json:"selectedEncodingMode"`
		PerTitleSettings      *StreamPerTitleSettings `json:"perTitleSettings"`
		Metadata              *StreamMetadata         `json:"metadata"`
		DecodingErrorMode     DecodingErrorMode       `json:"decodingErrorMode"`
		AppliedSettings       *AppliedStreamSettings  `json:"appliedSettings"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result Stream

	result.Name = data.Name
	result.Description = data.Description
	result.CreatedAt = data.CreatedAt
	result.ModifiedAt = data.ModifiedAt
	result.CustomData = data.CustomData
	result.Id = data.Id
	result.InputStreams = data.InputStreams
	result.Outputs = data.Outputs
	result.CreateQualityMetaData = data.CreateQualityMetaData
	result.CodecConfigId = data.CodecConfigId
	result.SegmentsEncoded = data.SegmentsEncoded
	result.IgnoredBy = data.IgnoredBy
	result.Mode = data.Mode
	result.SelectedEncodingMode = data.SelectedEncodingMode
	result.PerTitleSettings = data.PerTitleSettings
	result.Metadata = data.Metadata
	result.DecodingErrorMode = data.DecodingErrorMode
	result.AppliedSettings = data.AppliedSettings

	allOfConditions, err := UnmarshalAbstractCondition(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Conditions = &allOfConditions

	*m = result

	return nil
}
