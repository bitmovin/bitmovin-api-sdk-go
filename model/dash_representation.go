package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// DashRepresentation model
type DashRepresentation interface {
	// DashRepresentationTypeDiscriminator returns the discriminator type of the polymorphic model
	DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator
}

// BaseDashRepresentation is the fallback type for the polymorphic model DashRepresentation.
type BaseDashRepresentation struct {
	// Id of the resource (required)
	Id                *string                             `json:"id"`
	TypeDiscriminator DashRepresentationTypeDiscriminator `json:"typeDiscriminator"`
}

func (m BaseDashRepresentation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return m.TypeDiscriminator
}

// UnmarshalDashRepresentationSlice unmarshals polymorphic slices of DashRepresentation
func UnmarshalDashRepresentationSlice(reader io.Reader, consumer bitutils.Consumer) ([]DashRepresentation, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []DashRepresentation
	for _, element := range elements {
		obj, err := unmarshalDashRepresentation(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDashRepresentation unmarshals polymorphic DashRepresentation
func UnmarshalDashRepresentation(reader io.Reader, consumer bitutils.Consumer) (DashRepresentation, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDashRepresentation(data, consumer)
}

func unmarshalDashRepresentation(data []byte, consumer bitutils.Consumer) (DashRepresentation, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the typeDiscriminator property.
	var baseType BaseDashRepresentation
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of typeDiscriminator is used to determine which type to create and unmarshal the data into
	switch baseType.DashRepresentationTypeDiscriminator() {
	case "DRM_FMP4":
		var result DashFmp4DrmRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "FMP4":
		var result DashFmp4Representation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "WEBM":
		var result DashWebmRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CMAF":
		var result DashCmafRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CHUNKED_TEXT":
		var result DashChunkedTextRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "MP4":
		var result DashMp4Representation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DRM_MP4":
		var result DashMp4DrmRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "PROGRESSIVE_WEBM":
		var result DashProgressiveWebmRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "VTT":
		var result DashVttRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SPRITE":
		var result SpriteRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "IMSC":
		var result DashImscRepresentation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CONTENT_PROTECTION":
		var result ContentProtection
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
