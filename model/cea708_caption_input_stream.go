package model

import (
    "encoding/json"
)

// Cea708CaptionInputStream model
type Cea708CaptionInputStream struct {
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
    Id *string `json:"id,omitempty"`
    // Id of the Input (required)
    InputId *string `json:"inputId,omitempty"`
    // Path to media file (required)
    InputPath *string `json:"inputPath,omitempty"`
    // The channel number of the subtitle on the respective stream position. Must not be smaller than 1 (required)
    Channel *int32 `json:"channel,omitempty"`
}

func (m Cea708CaptionInputStream) InputStreamType() InputStreamType {
    return InputStreamType_CAPTION_CEA708
}
func (m Cea708CaptionInputStream) MarshalJSON() ([]byte, error) {
    type M Cea708CaptionInputStream
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "CAPTION_CEA708"

    return json.Marshal(x)
}


