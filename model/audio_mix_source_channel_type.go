package model
// AudioMixSourceChannelType : Type of this source channel
type AudioMixSourceChannelType string

// List of AudioMixSourceChannelType
const (
	AudioMixSourceChannelType_CHANNEL_NUMBER AudioMixSourceChannelType = "CHANNEL_NUMBER"
	AudioMixSourceChannelType_FRONT_LEFT AudioMixSourceChannelType = "FRONT_LEFT"
	AudioMixSourceChannelType_FRONT_RIGHT AudioMixSourceChannelType = "FRONT_RIGHT"
	AudioMixSourceChannelType_CENTER AudioMixSourceChannelType = "CENTER"
	AudioMixSourceChannelType_LOW_FREQUENCY AudioMixSourceChannelType = "LOW_FREQUENCY"
	AudioMixSourceChannelType_BACK_LEFT AudioMixSourceChannelType = "BACK_LEFT"
	AudioMixSourceChannelType_BACK_RIGHT AudioMixSourceChannelType = "BACK_RIGHT"
	AudioMixSourceChannelType_SURROUND_LEFT AudioMixSourceChannelType = "SURROUND_LEFT"
	AudioMixSourceChannelType_SURROUND_RIGHT AudioMixSourceChannelType = "SURROUND_RIGHT"
)