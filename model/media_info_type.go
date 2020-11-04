package model

// MediaInfoType : MediaInfoType model
type MediaInfoType string

// List of possible MediaInfoType values
const (
	MediaInfoType_AUDIO           MediaInfoType = "AUDIO"
	MediaInfoType_VIDEO           MediaInfoType = "VIDEO"
	MediaInfoType_SUBTITLES       MediaInfoType = "SUBTITLES"
	MediaInfoType_CLOSED_CAPTIONS MediaInfoType = "CLOSED_CAPTIONS"
	MediaInfoType_VTT             MediaInfoType = "VTT"
)
