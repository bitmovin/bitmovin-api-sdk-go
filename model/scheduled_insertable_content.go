package model

// ScheduledInsertableContent model
type ScheduledInsertableContent struct {
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
	// Id of the insertable content to play instead of the live stream
	ContentId *string `json:"contentId,omitempty"`
	// Time to to play the content in UTC: YYYY-MM-DDThh:mm:ssZ, if this property is not set the content will be played as soon as possible.
	RunAt *DateTime `json:"runAt,omitempty"`
	// Duration for how long to play the content. Cut off if shorter, loop if longer than actual duration. This property is required if the insertable content is an image.
	DurationInSeconds *float64 `json:"durationInSeconds,omitempty"`
	// Status of the scheduled insertable content.
	Status ScheduledInsertableContentStatus `json:"status,omitempty"`
}
