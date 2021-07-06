package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
	"io/ioutil"
)

// CodecConfiguration model
type CodecConfiguration interface {
	// CodecConfigType returns the discriminator type of the polymorphic model
	CodecConfigType() CodecConfigType
}

// BaseCodecConfiguration is the fallback type for the polymorphic model CodecConfiguration.
type BaseCodecConfiguration struct {
	// Name of the resource. Can be freely chosen by the user. (required)
	Name *string `json:"name"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id   *string         `json:"id"`
	Type CodecConfigType `json:"type"`
}

func (m BaseCodecConfiguration) CodecConfigType() CodecConfigType {
	return m.Type
}

// UnmarshalCodecConfigurationSlice unmarshals polymorphic slices of CodecConfiguration
func UnmarshalCodecConfigurationSlice(reader io.Reader, consumer bitutils.Consumer) ([]CodecConfiguration, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []CodecConfiguration
	for _, element := range elements {
		obj, err := unmarshalCodecConfiguration(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCodecConfiguration unmarshals polymorphic CodecConfiguration
func UnmarshalCodecConfiguration(reader io.Reader, consumer bitutils.Consumer) (CodecConfiguration, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCodecConfiguration(data, consumer)
}

func unmarshalCodecConfiguration(data []byte, consumer bitutils.Consumer) (CodecConfiguration, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var baseType BaseCodecConfiguration
	if err := consumer.Consume(buf, &baseType); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch baseType.CodecConfigType() {
	case "AAC":
		var result AacAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DTS_PASSTHROUGH":
		var result DtsPassthroughAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DVB_SUBTITLE":
		var result DvbSubtitleConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "HE_AAC_V1":
		var result HeAacV1AudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "HE_AAC_V2":
		var result HeAacV2AudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "H264":
		var result H264VideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "H265":
		var result H265VideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "VP9":
		var result Vp9VideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "VP8":
		var result Vp8VideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "MP2":
		var result Mp2AudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "MP3":
		var result Mp3AudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AC3":
		var result Ac3AudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DD":
		var result DolbyDigitalAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "EAC3":
		var result Eac3AudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DDPLUS":
		var result DolbyDigitalPlusAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "OPUS":
		var result OpusAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "VORBIS":
		var result VorbisAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "MJPEG":
		var result MjpegVideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "AV1":
		var result Av1VideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "DOLBY_ATMOS":
		var result DolbyAtmosAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "H262":
		var result H262VideoConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "PCM":
		var result PcmAudioConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	case "WEBVTT":
		var result WebVttConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return baseType, nil
	}
}
