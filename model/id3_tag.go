package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
    "io/ioutil"
)


// Id3Tag model
type Id3Tag interface {
    // Id3TagType returns the discriminator type of the polymorphic model
    Id3TagType() Id3TagType
}

// BaseId3Tag is the fallback type for the polymorphic model Id3Tag.
type BaseId3Tag struct {
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
    Id *string `json:"id"`
    PositionMode Id3TagPositionMode `json:"positionMode,omitempty"`
    // Frame number at which the Tag should be inserted
    Frame *int64 `json:"frame,omitempty"`
    // Time in seconds where the Tag should be inserted
    Time *float64 `json:"time,omitempty"`
    Type Id3TagType `json:"type"`
}

func (m BaseId3Tag) Id3TagType() Id3TagType {
    return m.Type
}

// UnmarshalId3TagSlice unmarshals polymorphic slices of Id3Tag
func UnmarshalId3TagSlice(reader io.Reader, consumer bitutils.Consumer) ([]Id3Tag, error) {
  var elements []json.RawMessage
  if err := consumer.Consume(reader, &elements); err != nil {
    return nil, err
  }

  var result []Id3Tag
  for _, element := range elements {
    obj, err := unmarshalId3Tag(element, consumer)
    if err != nil {
      return nil, err
    }
    result = append(result, obj)
  }
  return result, nil
}

// UnmarshalId3Tag unmarshals polymorphic Id3Tag
func UnmarshalId3Tag(reader io.Reader, consumer bitutils.Consumer) (Id3Tag, error) {
  // we need to read this twice, so first into a buffer
  data, err := ioutil.ReadAll(reader)
  if err != nil {
    return nil, err
  }
  return unmarshalId3Tag(data, consumer)
}

func unmarshalId3Tag(data []byte, consumer bitutils.Consumer) (Id3Tag, error) {
  buf := bytes.NewBuffer(data)
  buf2 := bytes.NewBuffer(data)

  // the first time this is read is to fetch the value of the type property.
  var baseType BaseId3Tag
  if err := consumer.Consume(buf, &baseType); err != nil {
    return nil, err
  }

  // The value of type is used to determine which type to create and unmarshal the data into
  switch baseType.Id3TagType() {
  case "RAW":
    var result RawId3Tag
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "FRAME_ID":
    var result FrameIdId3Tag
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil

  case "PLAIN_TEXT":
    var result PlaintextId3Tag
    if err := consumer.Consume(buf2, &result); err != nil {
      return nil, err
    }
    return result, nil


  default:
    return baseType, nil
  }
}


