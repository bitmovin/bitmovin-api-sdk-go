package model

// AgeRange : AgeRange model
type AgeRange string

// List of possible AgeRange values
const (
	AgeRange_CHILD        AgeRange = "CHILD"
	AgeRange_TEEN         AgeRange = "TEEN"
	AgeRange_TWENTIES     AgeRange = "TWENTIES"
	AgeRange_THIRTIES     AgeRange = "THIRTIES"
	AgeRange_FORTIES      AgeRange = "FORTIES"
	AgeRange_FIFTIES      AgeRange = "FIFTIES"
	AgeRange_SIXTIES_PLUS AgeRange = "SIXTIES_PLUS"
	AgeRange_UNKNOWN      AgeRange = "UNKNOWN"
)
