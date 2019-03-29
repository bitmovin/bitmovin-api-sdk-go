package model

type InternalChunkLength struct {
	// Defines how the internal chunk length for encoding will be determined
	Mode ChunkLengthMode `json:"mode,omitempty"`
	// Defines a custom internal chunk length in seconds to be used for encoding if mode is set to `CUSTOM`. Valid range is from 1 to 180 seconds
	CustomChunkLength *float64 `json:"customChunkLength,omitempty"`
}

