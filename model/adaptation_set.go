package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// AdaptationSet model
type AdaptationSet interface {
	// AdaptationSetType returns the discriminator type of the polymorphic model
	AdaptationSetType() AdaptationSetType
}

// BaseAdaptationSet is the fallback type for the polymorphic model AdaptationSet.
type BaseAdaptationSet struct {
	// Id of the resource (required)
	Id *string `json:"id"`
	// Custom adaptation set attributes
	CustomAttributes []CustomAttribute `json:"customAttributes,omitempty"`
	// Roles of the adaptation set
	Roles []AdaptationSetRole `json:"roles,omitempty"`
	// Provide signaling of CEA 607 and CEA 708
	Accessibilities []Accessibility `json:"accessibilities,omitempty"`
	// List of labels
	Labels []Label           `json:"labels,omitempty"`
	Type   AdaptationSetType `json:"type"`
}

func (m BaseAdaptationSet) AdaptationSetType() AdaptationSetType {
	return m.Type
}

// UnmarshalAdaptationSetSlice unmarshals polymorphic slices of AdaptationSet
func UnmarshalAdaptationSetSlice(reader io.Reader, consumer bitutils.Consumer) ([]AdaptationSet, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AdaptationSet
	for _, element := range elements {
		obj, err := unmarshalAdaptationSet(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAdaptationSet unmarshals polymorphic AdaptationSet
func UnmarshalAdaptationSet(reader io.Reader, consumer bitutils.Consumer) (AdaptationSet, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAdaptationSet(data, consumer)
}

func unmarshalAdaptationSet(data []byte, consumer bitutils.Consumer) (AdaptationSet, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseAdaptationSet
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.AdaptationSetType() {
	case "VIDEO":
		var result VideoAdaptationSet
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AUDIO":
		var result AudioAdaptationSet
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "IMAGE":
		var result ImageAdaptationSet
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SUBTITLE":
		var result SubtitleAdaptationSet
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
