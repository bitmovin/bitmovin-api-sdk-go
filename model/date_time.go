package model

import (
	"encoding/json"
	"time"
)

const (
	// RFC3339FullDate represents a full-date as specified by RFC3339
	// See: http://goo.gl/xXOvVd
	RFC3339FullDate = "2006-01-02"
	// RFC3339Millis represents a ISO8601 format to millis instead of to nanos
	RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"
	// RFC3339Micro represents a ISO8601 format to micro instead of to nano
	RFC3339Micro = "2006-01-02T15:04:05.000000Z07:00"
	// ISO8601LocalTime represents a ISO8601 format to ISO8601 in local time (no timezone)
	ISO8601LocalTime = "2006-01-02T15:04:05"
	ISO8601UTC       = "2006-01-02T15:04:05Z"

	jsonNull = "null"
)

var (
	dateTimeFormats = []string{RFC3339Micro, RFC3339Millis, time.RFC3339, time.RFC3339Nano, ISO8601LocalTime, ISO8601UTC}
)

// Date represents a date from the API
type Date time.Time

// NewDate constructs a new instance of type Date
func NewDate(year int, month time.Month, day int) Date {
	return Date(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// String converts this date into a string
func (d Date) String() string {
	return time.Time(d).Format(RFC3339FullDate)
}

// MarshalJSON returns the Date as JSON
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(RFC3339FullDate))
}

// UnmarshalJSON sets the Date from JSON
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}
	var strdate string
	if err := json.Unmarshal(data, &strdate); err != nil {
		return err
	}
	tt, err := time.Parse(RFC3339FullDate, strdate)
	if err != nil {
		return err
	}
	*d = Date(tt)
	return nil
}

// DateTime is a time but it serializes to ISO8601 format with millis
// It knows how to read 3 different variations of a RFC3339 date time.
// Most APIs we encounter want either millisecond or second precision times.
// This just tries to make it worry-free.
type DateTime time.Time

// NewDateTime is a representation of zero value for DateTime type
func NewDateTime() DateTime {
	return DateTime(time.Unix(0, 0).UTC())
}

// String converts this time to a string
func (t DateTime) String() string {
	return time.Time(t).Format(RFC3339Millis)
}

func (t DateTime) StringUTC() string {
	return time.Time(t).Format(ISO8601UTC)
}

// MarshalJSON returns the DateTime as JSON
func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(RFC3339Millis))
}

// UnmarshalJSON sets the DateTime from JSON
func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var tstr string
	if err := json.Unmarshal(data, &tstr); err != nil {
		return err
	}
	tt, err := ParseDateTime(tstr)
	if err != nil {
		return err
	}
	*t = tt
	return nil
}

// ParseDateTime parses a string that represents an ISO8601 time or a unix epoch
func ParseDateTime(data string) (DateTime, error) {
	if data == "" {
		return NewDateTime(), nil
	}
	var lastError error
	for _, layout := range dateTimeFormats {
		dd, err := time.Parse(layout, data)
		if err != nil {
			lastError = err
			continue
		}
		return DateTime(dd), nil
	}
	return DateTime{}, lastError
}
