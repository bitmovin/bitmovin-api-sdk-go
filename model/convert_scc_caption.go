package model


// ConvertSccCaption model
type ConvertSccCaption struct {
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
    // The input location to get the scc file from (required)
    Input *InputPath `json:"input,omitempty"`
    Outputs []EncodingOutput `json:"outputs,omitempty"`
    // Name of the captions file (required)
    FileName *string `json:"fileName,omitempty"`
    OutputFormat StreamCaptionOutputFormat `json:"outputFormat,omitempty"`
    // Optional settings when converting SCC to WebVTT
    WebVttSettings *ConvertSccCaptionWebVttSettings `json:"webVttSettings,omitempty"`
}



