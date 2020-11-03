package model


// CustomTag model
type CustomTag struct {
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
    // The positioning mode that should be used when inserting the placement opportunity (required)
    PositionMode PositionMode `json:"positionMode,omitempty"`
    // Id of keyframe where the custom tag should be inserted. Required, when KEYFRAME is selected as position mode.
    KeyframeId *string `json:"keyframeId,omitempty"`
    // Time in seconds where the custom tag should be inserted. Required, when TIME is selected as position mode.
    Time *float64 `json:"time,omitempty"`
    // The custom tag will be inserted before the specified segment. Required, when SEGMENT is selected as position mode.
    Segment *int64 `json:"segment,omitempty"`
    // The data to be contained in the custom tag. (required)
    Data *string `json:"data,omitempty"`
}



