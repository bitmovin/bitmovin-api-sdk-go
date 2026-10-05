package model

// Cea608ChannelType : CEA-608 caption channel, as defined in ANSI/CTA-608-E. Only the primary channel of each field is selectable.
type Cea608ChannelType string

// List of possible Cea608ChannelType values
const (
	Cea608ChannelType_CC1 Cea608ChannelType = "CC1"
	Cea608ChannelType_CC3 Cea608ChannelType = "CC3"
)
