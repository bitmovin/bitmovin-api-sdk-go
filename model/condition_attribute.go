package model

// ConditionAttribute : The attribute that should be checked
type ConditionAttribute string

// List of possible ConditionAttribute values
const (
	ConditionAttribute_HEIGHT                         ConditionAttribute = "HEIGHT"
	ConditionAttribute_WIDTH                          ConditionAttribute = "WIDTH"
	ConditionAttribute_BITRATE                        ConditionAttribute = "BITRATE"
	ConditionAttribute_FPS                            ConditionAttribute = "FPS"
	ConditionAttribute_ASPECTRATIO                    ConditionAttribute = "ASPECTRATIO"
	ConditionAttribute_INPUTSTREAM                    ConditionAttribute = "INPUTSTREAM"
	ConditionAttribute_LANGUAGE                       ConditionAttribute = "LANGUAGE"
	ConditionAttribute_CHANNELFORMAT                  ConditionAttribute = "CHANNELFORMAT"
	ConditionAttribute_CHANNELLAYOUT                  ConditionAttribute = "CHANNELLAYOUT"
	ConditionAttribute_STREAMCOUNT                    ConditionAttribute = "STREAMCOUNT"
	ConditionAttribute_AUDIOSTREAMCOUNT               ConditionAttribute = "AUDIOSTREAMCOUNT"
	ConditionAttribute_VIDEOSTREAMCOUNT               ConditionAttribute = "VIDEOSTREAMCOUNT"
	ConditionAttribute_DURATION                       ConditionAttribute = "DURATION"
	ConditionAttribute_ROTATION                       ConditionAttribute = "ROTATION"
	ConditionAttribute_CONNECTION_STATUS              ConditionAttribute = "CONNECTION_STATUS"
	ConditionAttribute_CONNECTION_STATUS_JUST_CHANGED ConditionAttribute = "CONNECTION_STATUS_JUST_CHANGED"
)
