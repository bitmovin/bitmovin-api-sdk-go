package bitutils

import (
	"encoding/json"
	"io"
)

// ConsumerFunc represents a function that can be used as a consumer
type ConsumerFunc func(io.Reader, interface{}) error

// Consume consumes the reader into the data parameter
func (fn ConsumerFunc) Consume(reader io.Reader, data interface{}) error {
	return fn(reader, data)
}

// Consumer implementations know how to bind the values on the provided interface to
// data provided by the request body
type Consumer interface {
	// Consume performs the binding of request values
	Consume(io.Reader, interface{}) error
}

func JSONConsumer() Consumer {
	return ConsumerFunc(func(reader io.Reader, data interface{}) error {
		dec := json.NewDecoder(reader)
		dec.UseNumber() // preserve number formats
		return dec.Decode(data)
	})
}
