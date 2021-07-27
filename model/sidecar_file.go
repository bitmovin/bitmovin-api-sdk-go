package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// A file that is added to an encoding. The size limit for a sidecar file is 10 MB
type SidecarFile interface {
	// SidecarFileType returns the discriminator type of the polymorphic model
	SidecarFileType() SidecarFileType
}

// BaseSidecarFile is the fallback type for the polymorphic model SidecarFile.
type BaseSidecarFile struct {
	// Id of the resource (required)
	Id *string `json:"id"`
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
	// Id of input (required)
	InputId *string `json:"inputId"`
	// Path to sidecar file (required)
	InputPath *string          `json:"inputPath"`
	Outputs   []EncodingOutput `json:"outputs,omitempty"`
	// This defines how errors should be handled
	ErrorMode SidecarErrorMode `json:"errorMode,omitempty"`
	Type      SidecarFileType  `json:"type"`
}

func (m BaseSidecarFile) SidecarFileType() SidecarFileType {
	return m.Type
}

// UnmarshalSidecarFileSlice unmarshals polymorphic slices of SidecarFile
func UnmarshalSidecarFileSlice(reader io.Reader, consumer bitutils.Consumer) ([]SidecarFile, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []SidecarFile
	for _, element := range elements {
		obj, err := unmarshalSidecarFile(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalSidecarFile unmarshals polymorphic SidecarFile
func UnmarshalSidecarFile(reader io.Reader, consumer bitutils.Consumer) (SidecarFile, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalSidecarFile(data, consumer)
}

func unmarshalSidecarFile(data []byte, consumer bitutils.Consumer) (SidecarFile, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseSidecarFile
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.SidecarFileType() {
	case "WEB_VTT":
		var result WebVttSidecarFile
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
