package model

// AnalyticsHttpRequestType : AnalyticsHttpRequestType model
type AnalyticsHttpRequestType string

// List of possible AnalyticsHttpRequestType values
const (
	AnalyticsHttpRequestType_DRM_LICENSE_WIDEVINE AnalyticsHttpRequestType = "DRM_LICENSE_WIDEVINE,"
	AnalyticsHttpRequestType_MEDIA_THUMBNAILS     AnalyticsHttpRequestType = "MEDIA_THUMBNAILS,"
	AnalyticsHttpRequestType_MEDIA_VIDEO          AnalyticsHttpRequestType = "MEDIA_VIDEO,"
	AnalyticsHttpRequestType_MEDIA_AUDIO          AnalyticsHttpRequestType = "MEDIA_AUDIO,"
	AnalyticsHttpRequestType_MEDIA_PROGRESSIVE    AnalyticsHttpRequestType = "MEDIA_PROGRESSIVE,"
	AnalyticsHttpRequestType_MEDIA_SUBTITLES      AnalyticsHttpRequestType = "MEDIA_SUBTITLES,"
	AnalyticsHttpRequestType_MANIFEST_DASH        AnalyticsHttpRequestType = "MANIFEST_DASH,"
	AnalyticsHttpRequestType_MANIFEST_HLS_MASTER  AnalyticsHttpRequestType = "MANIFEST_HLS_MASTER,"
	AnalyticsHttpRequestType_MANIFEST_HLS_VARIANT AnalyticsHttpRequestType = "MANIFEST_HLS_VARIANT,"
	AnalyticsHttpRequestType_MANIFEST_SMOOTH      AnalyticsHttpRequestType = "MANIFEST_SMOOTH,"
	AnalyticsHttpRequestType_KEY_HLS_AES          AnalyticsHttpRequestType = "KEY_HLS_AES,"
	AnalyticsHttpRequestType_UNKNOWN              AnalyticsHttpRequestType = "UNKNOWN"
)
