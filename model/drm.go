package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// Drm model
type Drm interface {
	// DrmType returns the discriminator type of the polymorphic model
	DrmType() DrmType
}

// BaseDrm is the fallback type for the polymorphic model Drm.
type BaseDrm struct {
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
	Outputs    []EncodingOutput        `json:"outputs,omitempty"`
	Type       DrmType                 `json:"type"`
}

func (m BaseDrm) DrmType() DrmType {
	return m.Type
}

// UnmarshalDrmSlice unmarshals polymorphic slices of Drm
func UnmarshalDrmSlice(reader io.Reader, consumer bitutils.Consumer) ([]Drm, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Drm
	for _, element := range elements {
		obj, err := unmarshalDrm(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDrm unmarshals polymorphic Drm
func UnmarshalDrm(reader io.Reader, consumer bitutils.Consumer) (Drm, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDrm(data, consumer)
}

func unmarshalDrm(data []byte, consumer bitutils.Consumer) (Drm, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseDrm
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.DrmType() {
	case "WIDEVINE":
		var result WidevineDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "PLAYREADY":
		var result PlayReadyDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "PRIMETIME":
		var result PrimeTimeDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "FAIRPLAY":
		var result FairPlayDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "MARLIN":
		var result MarlinDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CLEARKEY":
		var result ClearKeyDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AES":
		var result AesEncryptionDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CENC":
		var result CencDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SPEKE":
		var result SpekeDrm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
