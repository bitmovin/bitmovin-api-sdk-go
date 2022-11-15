package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// SimpleEncodingVodJobOutput model
type SimpleEncodingVodJobOutput interface {
	// SimpleEncodingVodJobOutputType returns the discriminator type of the polymorphic model
	SimpleEncodingVodJobOutputType() SimpleEncodingVodJobOutputType
}

// BaseSimpleEncodingVodJobOutput is the fallback type for the polymorphic model SimpleEncodingVodJobOutput.
type BaseSimpleEncodingVodJobOutput struct {
	// List of artifacts created by the encoding job. Artifacts are files essential for playback of the generated content, e.g. manifests.
	Artifacts []SimpleEncodingVodJobOutputArtifact `json:"artifacts,omitempty"`
	Type      SimpleEncodingVodJobOutputType       `json:"type"`
}

func (m BaseSimpleEncodingVodJobOutput) SimpleEncodingVodJobOutputType() SimpleEncodingVodJobOutputType {
	return m.Type
}

// UnmarshalSimpleEncodingVodJobOutputSlice unmarshals polymorphic slices of SimpleEncodingVodJobOutput
func UnmarshalSimpleEncodingVodJobOutputSlice(reader io.Reader, consumer bitutils.Consumer) ([]SimpleEncodingVodJobOutput, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []SimpleEncodingVodJobOutput
	for _, element := range elements {
		obj, err := unmarshalSimpleEncodingVodJobOutput(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalSimpleEncodingVodJobOutput unmarshals polymorphic SimpleEncodingVodJobOutput
func UnmarshalSimpleEncodingVodJobOutput(reader io.Reader, consumer bitutils.Consumer) (SimpleEncodingVodJobOutput, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalSimpleEncodingVodJobOutput(data, consumer)
}

func unmarshalSimpleEncodingVodJobOutput(data []byte, consumer bitutils.Consumer) (SimpleEncodingVodJobOutput, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseSimpleEncodingVodJobOutput
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.SimpleEncodingVodJobOutputType() {
	case "URL":
		var result SimpleEncodingVodJobUrlOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CDN":
		var result SimpleEncodingVodJobCdnOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
