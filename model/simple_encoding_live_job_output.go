package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// SimpleEncodingLiveJobOutput model
type SimpleEncodingLiveJobOutput interface {
	// SimpleEncodingLiveJobOutputType returns the discriminator type of the polymorphic model
	SimpleEncodingLiveJobOutputType() SimpleEncodingLiveJobOutputType
}

// BaseSimpleEncodingLiveJobOutput is the fallback type for the polymorphic model SimpleEncodingLiveJobOutput.
type BaseSimpleEncodingLiveJobOutput struct {
	// The maximum output resolution to be generated
	MaxResolution SimpleEncodingLiveMaxResolution `json:"maxResolution,omitempty"`
	Type          SimpleEncodingLiveJobOutputType `json:"type"`
}

func (m BaseSimpleEncodingLiveJobOutput) SimpleEncodingLiveJobOutputType() SimpleEncodingLiveJobOutputType {
	return m.Type
}

// UnmarshalSimpleEncodingLiveJobOutputSlice unmarshals polymorphic slices of SimpleEncodingLiveJobOutput
func UnmarshalSimpleEncodingLiveJobOutputSlice(reader io.Reader, consumer bitutils.Consumer) ([]SimpleEncodingLiveJobOutput, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []SimpleEncodingLiveJobOutput
	for _, element := range elements {
		obj, err := unmarshalSimpleEncodingLiveJobOutput(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalSimpleEncodingLiveJobOutput unmarshals polymorphic SimpleEncodingLiveJobOutput
func UnmarshalSimpleEncodingLiveJobOutput(reader io.Reader, consumer bitutils.Consumer) (SimpleEncodingLiveJobOutput, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalSimpleEncodingLiveJobOutput(data, consumer)
}

func unmarshalSimpleEncodingLiveJobOutput(data []byte, consumer bitutils.Consumer) (SimpleEncodingLiveJobOutput, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseSimpleEncodingLiveJobOutput
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.SimpleEncodingLiveJobOutputType() {
	case "URL":
		var result SimpleEncodingLiveJobUrlOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CDN":
		var result SimpleEncodingLiveJobCdnOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
