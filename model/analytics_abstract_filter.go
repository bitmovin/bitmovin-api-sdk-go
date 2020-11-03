package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
    "io/ioutil"
)


// AnalyticsAbstractFilter model
type AnalyticsAbstractFilter interface {
    // AnalyticsQueryOperator returns the discriminator type of the polymorphic model
    AnalyticsQueryOperator() AnalyticsQueryOperator
}

// BaseAnalyticsAbstractFilter is the fallback type for the polymorphic model AnalyticsAbstractFilter.
type BaseAnalyticsAbstractFilter struct {
    Name AnalyticsAttribute `json:"name"`
    Operator AnalyticsQueryOperator `json:"operator"`
}

func (m BaseAnalyticsAbstractFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return m.Operator
}

// UnmarshalAnalyticsAbstractFilterSlice unmarshals polymorphic slices of AnalyticsAbstractFilter
func UnmarshalAnalyticsAbstractFilterSlice(reader io.Reader, consumer bitutils.Consumer) ([]AnalyticsAbstractFilter, error) {
  var elements []json.RawMessage
  if err := consumer.Consume(reader, &elements); err != nil {
    return nil, err
  }

  var result []AnalyticsAbstractFilter
  for _, element := range elements {
    obj, err := unmarshalAnalyticsAbstractFilter(element, consumer)
    if err != nil {
      return nil, err
    }
    result = append(result, obj)
  }
  return result, nil
}

// UnmarshalAnalyticsAbstractFilter unmarshals polymorphic AnalyticsAbstractFilter
func UnmarshalAnalyticsAbstractFilter(reader io.Reader, consumer bitutils.Consumer) (AnalyticsAbstractFilter, error) {
  // we need to read this twice, so first into a buffer
  data, err := ioutil.ReadAll(reader)
  if err != nil {
    return nil, err
  }
  return unmarshalAnalyticsAbstractFilter(data, consumer)
}

func unmarshalAnalyticsAbstractFilter(data []byte, consumer bitutils.Consumer) (AnalyticsAbstractFilter, error) {
  buf := bytes.NewBuffer(data)
  buf2 := bytes.NewBuffer(data)

  // the first time this is read is to fetch the value of the operator property.
  var baseType BaseAnalyticsAbstractFilter
  if err := consumer.Consume(buf, &baseType); err != nil {
    return nil, err
  }

  // The value of operator is used to determine which type to create and unmarshal the data into
  switch baseType.AnalyticsQueryOperator() {
  case "IN":
    var result AnalyticsInFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "EQ":
    var result AnalyticsEqualFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "NE":
    var result AnalyticsNotEqualFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "LT":
    var result AnalyticsLessThanFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "LTE":
    var result AnalyticsLessThanOrEqualFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "GT":
    var result AnalyticsGreaterThanFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "GTE":
    var result AnalyticsGreaterThanOrEqualFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "CONTAINS":
    var result AnalyticsContainsFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "NOTCONTAINS":
    var result AnalyticsNotContainsFilter
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil


  default:
    return baseType, nil
  }
}


