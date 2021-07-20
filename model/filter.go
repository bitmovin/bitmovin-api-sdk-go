package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// Filter model
type Filter interface {
	// FilterType returns the discriminator type of the polymorphic model
	FilterType() FilterType
}

// BaseFilter is the fallback type for the polymorphic model Filter.
type BaseFilter struct {
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
	Type       FilterType              `json:"type"`
}

func (m BaseFilter) FilterType() FilterType {
	return m.Type
}

// UnmarshalFilterSlice unmarshals polymorphic slices of Filter
func UnmarshalFilterSlice(reader io.Reader, consumer bitutils.Consumer) ([]Filter, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Filter
	for _, element := range elements {
		obj, err := unmarshalFilter(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalFilter unmarshals polymorphic Filter
func UnmarshalFilter(reader io.Reader, consumer bitutils.Consumer) (Filter, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalFilter(data, consumer)
}

func unmarshalFilter(data []byte, consumer bitutils.Consumer) (Filter, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseFilter
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.FilterType() {
	case "CROP":
		var result CropFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CONFORM":
		var result ConformFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "WATERMARK":
		var result WatermarkFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "ENHANCED_WATERMARK":
		var result EnhancedWatermarkFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "ROTATE":
		var result RotateFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DEINTERLACE":
		var result DeinterlaceFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "ENHANCED_DEINTERLACE":
		var result EnhancedDeinterlaceFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AUDIO_MIX":
		var result AudioMixFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DENOISE_HQDN3D":
		var result DenoiseHqdn3dFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "TEXT":
		var result TextFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "UNSHARP":
		var result UnsharpFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SCALE":
		var result ScaleFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "INTERLACE":
		var result InterlaceFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AUDIO_VOLUME":
		var result AudioVolumeFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "EBU_R128_SINGLE_PASS":
		var result EbuR128SinglePassFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
