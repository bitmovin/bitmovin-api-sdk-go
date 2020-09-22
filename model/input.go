package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// Input model
type Input interface {
	// InputType returns the discriminator type of the polymorphic model
	InputType() InputType
}

// BaseInput is the fallback type for the polymorphic model Input.
type BaseInput struct {
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
	// Id of the resource (required)
	Id   *string   `json:"id"`
	Type InputType `json:"type"`
}

func (m BaseInput) InputType() InputType {
	return m.Type
}

// UnmarshalInputSlice unmarshals polymorphic slices of Input
func UnmarshalInputSlice(reader io.Reader, consumer bitutils.Consumer) ([]Input, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Input
	for _, element := range elements {
		obj, err := unmarshalInput(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalInput unmarshals polymorphic Input
func UnmarshalInput(reader io.Reader, consumer bitutils.Consumer) (Input, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalInput(data, consumer)
}

func unmarshalInput(data []byte, consumer bitutils.Consumer) (Input, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseInput
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.InputType() {
	case "AKAMAI_NETSTORAGE":
		var result AkamaiNetStorageInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "ASPERA":
		var result AsperaInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AZURE":
		var result AzureInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "REDUNDANT_RTMP":
		var result RedundantRtmpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "FTP":
		var result FtpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GENERIC_S3":
		var result GenericS3Input
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GCS":
		var result GcsInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "HTTP":
		var result HttpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "HTTPS":
		var result HttpsInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "LOCAL":
		var result LocalInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "RTMP":
		var result RtmpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "S3":
		var result S3Input
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "S3_ROLE_BASED":
		var result S3RoleBasedInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SFTP":
		var result SftpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "TCP":
		var result TcpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "UDP":
		var result UdpInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "UDP_MULTICAST":
		var result UdpMulticastInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "ZIXI":
		var result ZixiInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SRT":
		var result SrtInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GCS_SERVICE_ACCOUNT":
		var result GcsServiceAccountInput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
