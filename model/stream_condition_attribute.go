package model
// StreamConditionAttribute : The attribute that should be checked
type StreamConditionAttribute string

// List of StreamConditionAttribute
const (
	StreamConditionAttribute_MEDIA_TYPE StreamConditionAttribute = "MEDIA_TYPE"
	StreamConditionAttribute_STREAM_ID StreamConditionAttribute = "STREAM_ID"
	StreamConditionAttribute_BITS_READ_AVG StreamConditionAttribute = "BITS_READ_AVG"
	StreamConditionAttribute_BITS_READ_MIN StreamConditionAttribute = "BITS_READ_MIN"
	StreamConditionAttribute_BITS_READ_MAX StreamConditionAttribute = "BITS_READ_MAX"
	StreamConditionAttribute_SAMPLES_READ_AVG StreamConditionAttribute = "SAMPLES_READ_AVG"
	StreamConditionAttribute_SAMPLES_READ_MIN StreamConditionAttribute = "SAMPLES_READ_MIN"
	StreamConditionAttribute_SAMPLES_READ_MAX StreamConditionAttribute = "SAMPLES_READ_MAX"
	StreamConditionAttribute_WIDTH StreamConditionAttribute = "WIDTH"
	StreamConditionAttribute_HEIGHT StreamConditionAttribute = "HEIGHT"
	StreamConditionAttribute_CODEC StreamConditionAttribute = "CODEC"
)