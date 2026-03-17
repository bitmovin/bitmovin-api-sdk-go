package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// AgentChatMessagePart model
type AgentChatMessagePart interface {
	// AgentChatMessagePartType returns the discriminator type of the polymorphic model
	AgentChatMessagePartType() AgentChatMessagePartType
}

// BaseAgentChatMessagePart is the fallback type for the polymorphic model AgentChatMessagePart.
type BaseAgentChatMessagePart struct {
	Type AgentChatMessagePartType `json:"type"`
}

func (m BaseAgentChatMessagePart) AgentChatMessagePartType() AgentChatMessagePartType {
	return m.Type
}

// UnmarshalAgentChatMessagePartSlice unmarshals polymorphic slices of AgentChatMessagePart
func UnmarshalAgentChatMessagePartSlice(reader io.Reader, consumer bitutils.Consumer) ([]AgentChatMessagePart, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AgentChatMessagePart
	for _, element := range elements {
		obj, err := unmarshalAgentChatMessagePart(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAgentChatMessagePart unmarshals polymorphic AgentChatMessagePart
func UnmarshalAgentChatMessagePart(reader io.Reader, consumer bitutils.Consumer) (AgentChatMessagePart, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAgentChatMessagePart(data, consumer)
}

func unmarshalAgentChatMessagePart(data []byte, consumer bitutils.Consumer) (AgentChatMessagePart, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseAgentChatMessagePart
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.AgentChatMessagePartType() {
	case "text":
		var result AgentChatTextPart
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "data-attachments":
		var result AgentChatAttachmentsPart
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "dynamic-tool":
		var result AgentChatDynamicToolPart
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
