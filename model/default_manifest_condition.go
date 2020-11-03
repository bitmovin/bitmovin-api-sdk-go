package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
    "io/ioutil"
)


// DefaultManifestCondition model
type DefaultManifestCondition interface {
    // ConditionType returns the discriminator type of the polymorphic model
    ConditionType() ConditionType
}

// BaseDefaultManifestCondition is the fallback type for the polymorphic model DefaultManifestCondition.
type BaseDefaultManifestCondition struct {
    Type ConditionType `json:"type"`
}

func (m BaseDefaultManifestCondition) ConditionType() ConditionType {
    return m.Type
}

// UnmarshalDefaultManifestConditionSlice unmarshals polymorphic slices of DefaultManifestCondition
func UnmarshalDefaultManifestConditionSlice(reader io.Reader, consumer bitutils.Consumer) ([]DefaultManifestCondition, error) {
  var elements []json.RawMessage
  if err := consumer.Consume(reader, &elements); err != nil {
    return nil, err
  }

  var result []DefaultManifestCondition
  for _, element := range elements {
    obj, err := unmarshalDefaultManifestCondition(element, consumer)
    if err != nil {
      return nil, err
    }
    result = append(result, obj)
  }
  return result, nil
}

// UnmarshalDefaultManifestCondition unmarshals polymorphic DefaultManifestCondition
func UnmarshalDefaultManifestCondition(reader io.Reader, consumer bitutils.Consumer) (DefaultManifestCondition, error) {
  // we need to read this twice, so first into a buffer
  data, err := ioutil.ReadAll(reader)
  if err != nil {
    return nil, err
  }
  return unmarshalDefaultManifestCondition(data, consumer)
}

func unmarshalDefaultManifestCondition(data []byte, consumer bitutils.Consumer) (DefaultManifestCondition, error) {
  buf := bytes.NewBuffer(data)
  buf2 := bytes.NewBuffer(data)

  // the first time this is read is to fetch the value of the type property.
  var baseType BaseDefaultManifestCondition
  if err := consumer.Consume(buf, &baseType); err != nil {
    return nil, err
  }

  // The value of type is used to determine which type to create and unmarshal the data into
  switch baseType.ConditionType() {
  case "CONDITION":
    var result DefaultManifestAttributeCondition
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "AND":
    var result DefaultManifestAndCondition
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "OR":
    var result DefaultManifestOrCondition
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil


  default:
    return baseType, nil
  }
}


