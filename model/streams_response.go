package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// StreamsResponse model
type StreamsResponse interface {
	// StreamsType returns the discriminator type of the polymorphic model
	StreamsType() StreamsType
}

// BaseStreamsResponse is the fallback type for the polymorphic model StreamsResponse.
type BaseStreamsResponse struct {
	// The identifier of the stream
	Id *string `json:"id,omitempty"`
	// The title of the stream
	Title *string `json:"title,omitempty"`
	// The description of the stream
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Type of the Stream contained in the StreamsResponse
	Type StreamsType `json:"type,omitempty"`
}

func (m BaseStreamsResponse) StreamsType() StreamsType {
	return m.Type
}

// UnmarshalStreamsResponseSlice unmarshals polymorphic slices of StreamsResponse
func UnmarshalStreamsResponseSlice(reader io.Reader, consumer bitutils.Consumer) ([]StreamsResponse, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []StreamsResponse
	for _, element := range elements {
		obj, err := unmarshalStreamsResponse(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalStreamsResponse unmarshals polymorphic StreamsResponse
func UnmarshalStreamsResponse(reader io.Reader, consumer bitutils.Consumer) (StreamsResponse, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalStreamsResponse(data, consumer)
}

func unmarshalStreamsResponse(data []byte, consumer bitutils.Consumer) (StreamsResponse, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the  property.
	var baseType BaseStreamsResponse
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of  is used to determine which type to create and unmarshal the data into
	switch baseType.StreamsType() {
	case "VIDEO":
		var result StreamsVideoResponse
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "LIVE":
		var result StreamsLiveResponse
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
