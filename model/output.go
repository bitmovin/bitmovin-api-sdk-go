package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// Output model
type Output interface {
	// OutputType returns the discriminator type of the polymorphic model
	OutputType() OutputType
}

// BaseOutput is the fallback type for the polymorphic model Output.
type BaseOutput struct {
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
	Id   *string    `json:"id"`
	Acl  []AclEntry `json:"acl,omitempty"`
	Type OutputType `json:"type"`
}

func (m BaseOutput) OutputType() OutputType {
	return m.Type
}

// UnmarshalOutputSlice unmarshals polymorphic slices of Output
func UnmarshalOutputSlice(reader io.Reader, consumer bitutils.Consumer) ([]Output, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Output
	for _, element := range elements {
		obj, err := unmarshalOutput(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalOutput unmarshals polymorphic Output
func UnmarshalOutput(reader io.Reader, consumer bitutils.Consumer) (Output, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalOutput(data, consumer)
}

func unmarshalOutput(data []byte, consumer bitutils.Consumer) (Output, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseOutput
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.OutputType() {
	case "AKAMAI_NETSTORAGE":
		var result AkamaiNetStorageOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AZURE":
		var result AzureOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GENERIC_S3":
		var result GenericS3Output
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GCS":
		var result GcsOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "FTP":
		var result FtpOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "LOCAL":
		var result LocalOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "S3":
		var result S3Output
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "S3_ROLE_BASED":
		var result S3RoleBasedOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "SFTP":
		var result SftpOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AKAMAI_MSL":
		var result AkamaiMslOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "LIVE_MEDIA_INGEST":
		var result LiveMediaIngestOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GCS_SERVICE_ACCOUNT":
		var result GcsServiceAccountOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
