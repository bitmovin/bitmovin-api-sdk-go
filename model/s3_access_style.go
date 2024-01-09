package model

// S3AccessStyle : Specifies whether to use path or virtual-hosted style access
type S3AccessStyle string

// List of possible S3AccessStyle values
const (
	S3AccessStyle_VIRTUAL_HOSTED S3AccessStyle = "VIRTUAL_HOSTED"
	S3AccessStyle_PATH           S3AccessStyle = "PATH"
)
