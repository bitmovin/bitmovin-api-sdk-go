package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
    "io/ioutil"
)


// AbstractCondition model
type AbstractCondition interface {
    // ConditionType returns the discriminator type of the polymorphic model
    ConditionType() ConditionType
}

// BaseAbstractCondition is the fallback type for the polymorphic model AbstractCondition.
type BaseAbstractCondition struct {
    Type ConditionType `json:"type"`
}

func (m BaseAbstractCondition) ConditionType() ConditionType {
    return m.Type
}

// UnmarshalAbstractConditionSlice unmarshals polymorphic slices of AbstractCondition
func UnmarshalAbstractConditionSlice(reader io.Reader, consumer bitutils.Consumer) ([]AbstractCondition, error) {
  var elements []json.RawMessage
  if err := consumer.Consume(reader, &elements); err != nil {
    return nil, err
  }

  var result []AbstractCondition
  for _, element := range elements {
    obj, err := unmarshalAbstractCondition(element, consumer)
    if err != nil {
      return nil, err
    }
    result = append(result, obj)
  }
  return result, nil
}

// UnmarshalAbstractCondition unmarshals polymorphic AbstractCondition
func UnmarshalAbstractCondition(reader io.Reader, consumer bitutils.Consumer) (AbstractCondition, error) {
  // we need to read this twice, so first into a buffer
  data, err := ioutil.ReadAll(reader)
  if err != nil {
    return nil, err
  }
  return unmarshalAbstractCondition(data, consumer)
}

func unmarshalAbstractCondition(data []byte, consumer bitutils.Consumer) (AbstractCondition, error) {
  buf := bytes.NewBuffer(data)
  buf2 := bytes.NewBuffer(data)

  // the first time this is read is to fetch the value of the type property.
  var baseType BaseAbstractCondition
  if err := consumer.Consume(buf, &baseType); err != nil {
    return nil, err
  }

  // The value of type is used to determine which type to create and unmarshal the data into
  switch baseType.ConditionType() {
  case "CONDITION":
    var result Condition
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "AND":
    var result AndConjunction
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "OR":
    var result OrConjunction
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil


  default:
    return baseType, nil
  }
}


