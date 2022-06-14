package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// SimpleEncodingVodJobInput model
type SimpleEncodingVodJobInput interface {
	// SimpleEncodingVodJobInputSourceType returns the discriminator type of the polymorphic model
	SimpleEncodingVodJobInputSourceType() SimpleEncodingVodJobInputSourceType
}

// BaseSimpleEncodingVodJobInput is the fallback type for the polymorphic model SimpleEncodingVodJobInput.
type BaseSimpleEncodingVodJobInput struct {
	Type SimpleEncodingVodJobInputSourceType `json:"type"`
}

func (m BaseSimpleEncodingVodJobInput) SimpleEncodingVodJobInputSourceType() SimpleEncodingVodJobInputSourceType {
	return m.Type
}

// UnmarshalSimpleEncodingVodJobInputSlice unmarshals polymorphic slices of SimpleEncodingVodJobInput
func UnmarshalSimpleEncodingVodJobInputSlice(reader io.Reader, consumer bitutils.Consumer) ([]SimpleEncodingVodJobInput, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []SimpleEncodingVodJobInput
	for _, element := range elements {
		obj, err := unmarshalSimpleEncodingVodJobInput(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalSimpleEncodingVodJobInput unmarshals polymorphic SimpleEncodingVodJobInput
func UnmarshalSimpleEncodingVodJobInput(reader io.Reader, consumer bitutils.Consumer) (SimpleEncodingVodJobInput, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalSimpleEncodingVodJobInput(data, consumer)
}

func unmarshalSimpleEncodingVodJobInput(data []byte, consumer bitutils.Consumer) (SimpleEncodingVodJobInput, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseSimpleEncodingVodJobInput
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.SimpleEncodingVodJobInputSourceType() {
	case "URL":
		var result SimpleEncodingVodJobUrlInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DIRECT_FILE_UPLOAD":
		var result SimpleEncodingVodJobDirectFileUploadInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
