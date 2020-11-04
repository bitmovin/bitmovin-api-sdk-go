package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// AdAnalyticsAbstractFilter model
type AdAnalyticsAbstractFilter interface {
	// AnalyticsQueryOperator returns the discriminator type of the polymorphic model
	AnalyticsQueryOperator() AnalyticsQueryOperator
}

// BaseAdAnalyticsAbstractFilter is the fallback type for the polymorphic model AdAnalyticsAbstractFilter.
type BaseAdAnalyticsAbstractFilter struct {
	Name     AdAnalyticsAttribute   `json:"name"`
	Operator AnalyticsQueryOperator `json:"operator"`
}

func (m BaseAdAnalyticsAbstractFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return m.Operator
}

// UnmarshalAdAnalyticsAbstractFilterSlice unmarshals polymorphic slices of AdAnalyticsAbstractFilter
func UnmarshalAdAnalyticsAbstractFilterSlice(reader io.Reader, consumer bitutils.Consumer) ([]AdAnalyticsAbstractFilter, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AdAnalyticsAbstractFilter
	for _, element := range elements {
		obj, err := unmarshalAdAnalyticsAbstractFilter(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAdAnalyticsAbstractFilter unmarshals polymorphic AdAnalyticsAbstractFilter
func UnmarshalAdAnalyticsAbstractFilter(reader io.Reader, consumer bitutils.Consumer) (AdAnalyticsAbstractFilter, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAdAnalyticsAbstractFilter(data, consumer)
}

func unmarshalAdAnalyticsAbstractFilter(data []byte, consumer bitutils.Consumer) (AdAnalyticsAbstractFilter, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the operator property.
	var baseType BaseAdAnalyticsAbstractFilter
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of operator is used to determine which type to create and unmarshal the data into
	switch baseType.AnalyticsQueryOperator() {
	case "IN":
		var result AdAnalyticsInFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "EQ":
		var result AdAnalyticsEqualFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "NE":
		var result AdAnalyticsNotEqualFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "LT":
		var result AdAnalyticsLessThanFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "LTE":
		var result AdAnalyticsLessThanOrEqualFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GT":
		var result AdAnalyticsGreaterThanFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "GTE":
		var result AdAnalyticsGreaterThanOrEqualFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "CONTAINS":
		var result AdAnalyticsContainsFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "NOTCONTAINS":
		var result AdAnalyticsNotContainsFilter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
