package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// Rate control mode configuration. Use H265V2PerceptualQualityModeConfig for PQP mode or H265V2ConstantBitrateModeConfig for CBR mode.
type H265V2RateControlModeConfig interface {
	// H265V2RateControlModeConfigType returns the discriminator type of the polymorphic model
	H265V2RateControlModeConfigType() H265V2RateControlModeConfigType
}

// BaseH265V2RateControlModeConfig is the fallback type for the polymorphic model H265V2RateControlModeConfig.
type BaseH265V2RateControlModeConfig struct {
	Type H265V2RateControlModeConfigType `json:"type,omitempty"`
}

func (m BaseH265V2RateControlModeConfig) H265V2RateControlModeConfigType() H265V2RateControlModeConfigType {
	return m.Type
}

// UnmarshalH265V2RateControlModeConfigSlice unmarshals polymorphic slices of H265V2RateControlModeConfig
func UnmarshalH265V2RateControlModeConfigSlice(reader io.Reader, consumer bitutils.Consumer) ([]H265V2RateControlModeConfig, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []H265V2RateControlModeConfig
	for _, element := range elements {
		obj, err := unmarshalH265V2RateControlModeConfig(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalH265V2RateControlModeConfig unmarshals polymorphic H265V2RateControlModeConfig
func UnmarshalH265V2RateControlModeConfig(reader io.Reader, consumer bitutils.Consumer) (H265V2RateControlModeConfig, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalH265V2RateControlModeConfig(data, consumer)
}

func unmarshalH265V2RateControlModeConfig(data []byte, consumer bitutils.Consumer) (H265V2RateControlModeConfig, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the  property.
	var baseType BaseH265V2RateControlModeConfig
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of  is used to determine which type to create and unmarshal the data into
	switch baseType.H265V2RateControlModeConfigType() {
	case "PERCEPTUAL_QUALITY_MODE":
		var result H265V2PerceptualQualityModeConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CONSTANT_BITRATE_MODE":
		var result H265V2ConstantBitrateModeConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
