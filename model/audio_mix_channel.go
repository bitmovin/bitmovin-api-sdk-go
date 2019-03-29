package model

type AudioMixChannel struct {
	// Channel number of this mix (starting with 0)
	ChannelNumber *int32 `json:"channelNumber,omitempty"`
	// List of source channels to be mixed
	SourceChannels []SourceChannel `json:"sourceChannels,omitempty"`
}

