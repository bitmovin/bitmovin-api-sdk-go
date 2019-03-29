package model
// ChunkLengthMode : Defines how the internal chunk length for encoding will be determined
type ChunkLengthMode string

// List of ChunkLengthMode
const (
	ChunkLengthMode_SPEED_OPTIMIZED ChunkLengthMode = "SPEED_OPTIMIZED"
	ChunkLengthMode_CUSTOM ChunkLengthMode = "CUSTOM"
)