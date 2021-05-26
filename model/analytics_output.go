package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// AnalyticsOutput model
type AnalyticsOutput interface {
	// AnalyticsOutputType returns the discriminator type of the polymorphic model
	AnalyticsOutputType() AnalyticsOutputType
}

// BaseAnalyticsOutput is the fallback type for the polymorphic model AnalyticsOutput.
type BaseAnalyticsOutput struct {
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
	Id   *string             `json:"id"`
	Acl  []AclEntry          `json:"acl,omitempty"`
	Type AnalyticsOutputType `json:"type"`
}

func (m BaseAnalyticsOutput) AnalyticsOutputType() AnalyticsOutputType {
	return m.Type
}

// UnmarshalAnalyticsOutputSlice unmarshals polymorphic slices of AnalyticsOutput
func UnmarshalAnalyticsOutputSlice(reader io.Reader, consumer bitutils.Consumer) ([]AnalyticsOutput, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AnalyticsOutput
	for _, element := range elements {
		obj, err := unmarshalAnalyticsOutput(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAnalyticsOutput unmarshals polymorphic AnalyticsOutput
func UnmarshalAnalyticsOutput(reader io.Reader, consumer bitutils.Consumer) (AnalyticsOutput, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAnalyticsOutput(data, consumer)
}

func unmarshalAnalyticsOutput(data []byte, consumer bitutils.Consumer) (AnalyticsOutput, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseAnalyticsOutput
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.AnalyticsOutputType() {
	case "S3_ROLE_BASED":
		var result AnalyticsS3RoleBasedOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GCS_SERVICE_ACCOUNT":
		var result AnalyticsGcsServiceAccountOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AZURE":
		var result AnalyticsAzureOutput
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
