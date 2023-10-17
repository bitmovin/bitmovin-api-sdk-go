package model

// SignupSource : platform which initiated the creation of the resource
type SignupSource string

// List of possible SignupSource values
const (
	SignupSource_AWS      SignupSource = "AWS"
	SignupSource_AZURE    SignupSource = "AZURE"
	SignupSource_BITMOVIN SignupSource = "BITMOVIN"
	SignupSource_GOOGLE   SignupSource = "GOOGLE"
)
