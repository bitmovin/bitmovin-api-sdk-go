package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// InputStream model
type InputStream interface {
	// InputStreamType returns the discriminator type of the polymorphic model
	InputStreamType() InputStreamType
}

// BaseInputStream is the fallback type for the polymorphic model InputStream.
type BaseInputStream struct {
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
	Type       InputStreamType         `json:"type"`
}

func (m BaseInputStream) InputStreamType() InputStreamType {
	return m.Type
}

// UnmarshalInputStreamSlice unmarshals polymorphic slices of InputStream
func UnmarshalInputStreamSlice(reader io.Reader, consumer bitutils.Consumer) ([]InputStream, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []InputStream
	for _, element := range elements {
		obj, err := unmarshalInputStream(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalInputStream unmarshals polymorphic InputStream
func UnmarshalInputStream(reader io.Reader, consumer bitutils.Consumer) (InputStream, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalInputStream(data, consumer)
}

func unmarshalInputStream(data []byte, consumer bitutils.Consumer) (InputStream, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseInputStream
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.InputStreamType() {
	case "INGEST":
		var result IngestInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CONCATENATION":
		var result ConcatenationInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "TRIMMING_TIME_BASED":
		var result TimeBasedTrimmingInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "TRIMMING_TIME_CODE_TRACK":
		var result TimecodeTrackTrimmingInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "TRIMMING_H264_PICTURE_TIMING":
		var result H264PictureTimingTrimmingInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AUDIO_MIX":
		var result AudioMixInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SIDECAR_DOLBY_VISION_METADATA":
		var result DolbyVisionMetadataIngestInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CAPTION_CEA608":
		var result Cea608CaptionInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CAPTION_CEA708":
		var result Cea708CaptionInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "FILE":
		var result FileInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DVB_SUBTITLE":
		var result DvbSubtitleInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DOLBY_ATMOS":
		var result DolbyAtmosIngestInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DOLBY_VISION":
		var result DolbyVisionInputStream
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
