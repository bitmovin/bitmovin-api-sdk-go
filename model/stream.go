package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// Stream model
type Stream struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
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
	// Determines the input source(s) for the stream. All video streams of an encoding need to have identical input configurations (required)
	InputStreams []StreamInput    `json:"inputStreams,omitempty"`
	Outputs      []EncodingOutput `json:"outputs,omitempty"`
	// Set true to create quality metadata for this stream
	CreateQualityMetaData *bool `json:"createQualityMetaData,omitempty"`
	// Id of the codec configuration (required)
	CodecConfigId *string `json:"codecConfigId,omitempty"`
	// Number of encoded segments. Available after encoding finishes.
	SegmentsEncoded *int32 `json:"segmentsEncoded,omitempty"`
	// Defines a condition that is evaluated against the input of the Stream. If the condition is not fulfilled, the Stream will be ignored during the encoding process. The 'streamConditionMode' of a Muxing allows to control how ignored Streams affect the Muxing. When retrieving the resource after the analysis step of the encoding has finished, 'ignoredBy' will indicate if and why it has been ignored. See [Stream Conditions](https://bitmovin.com/docs/encoding/articles/stream-conditions) for more information
	Conditions *AbstractCondition `json:"conditions,omitempty"`
	// This read-only property is set during the analysis step of the encoding. If it contains items, the Stream has been ignored during the encoding process due to its 'conditions'
	IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
	// Mode of the stream
	Mode StreamMode `json:"mode,omitempty"`
	// The encoding mode that was used for this stream. This is derived from `encodingMode`, which can be specified in the codec configuration or in the encoding start request. Note that all streams of an encoding need to use the same encoding mode. This will therefore always match `selectedEncodingMode` of the related Encoding resource. Especially useful when `encodingMode` was not set explicitly or set to STANDARD (which translates to one of the other possible values on encoding start).
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
		Id                    *string                 `json:"id"`
		Name                  *string                 `json:"name"`
		Description           *string                 `json:"description"`
		CreatedAt             *DateTime               `json:"createdAt"`
		ModifiedAt            *DateTime               `json:"modifiedAt"`
		CustomData            *map[string]interface{} `json:"customData"`
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

	result.Id = data.Id
	result.Name = data.Name
	result.Description = data.Description
	result.CreatedAt = data.CreatedAt
	result.ModifiedAt = data.ModifiedAt
	result.CustomData = data.CustomData
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
