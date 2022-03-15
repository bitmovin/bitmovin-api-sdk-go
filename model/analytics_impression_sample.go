package model

// AnalyticsImpressionSample model
type AnalyticsImpressionSample struct {
	// Is an ad playing. 0 indicates no, 1 indicates yes
	Ad *int64 `json:"ad,omitempty"`
	// Collector version
	AnalyticsVersion *string `json:"analytics_version,omitempty"`
	// Audio Bitrate
	AudioBitrate *int64 `json:"audio_bitrate,omitempty"`
	// Audio codec of currently playing stream
	AudioCodec *string `json:"audio_codec,omitempty"`
	// Selected audio language
	AudioLanguage *string `json:"audio_language,omitempty"`
	// Autoplay enabled
	Autoplay *bool `json:"autoplay,omitempty"`
	// Browser name
	Browser *string `json:"browser,omitempty"`
	// Browser version major
	BrowserVersionMajor *string `json:"browser_version_major,omitempty"`
	// Browser version minor
	BrowserVersionMinor *string `json:"browser_version_minor,omitempty"`
	// Milliseconds the player buffered
	Buffered *int64 `json:"buffered,omitempty"`
	// CDN Provider
	CdnProvider *string `json:"cdn_provider,omitempty"`
	// Casting Technology
	CastTech *string `json:"cast_tech,omitempty"`
	// City
	City *string `json:"city,omitempty"`
	// Current time of the client
	ClientTime *int64 `json:"client_time,omitempty"`
	// Country
	Country *string `json:"country,omitempty"`
	// Custom user ID
	CustomUserId *string `json:"custom_user_id,omitempty"`
	// Free form data set via the customData1 field in the analytics collector configuration
	CustomData1 *string `json:"custom_data_1,omitempty"`
	// Free form data set via the customData2 field in the analytics collector configuration
	CustomData2 *string `json:"custom_data_2,omitempty"`
	// Free form data set via the customData3 field in the analytics collector configuration
	CustomData3 *string `json:"custom_data_3,omitempty"`
	// Free form data set via the customData4 field in the analytics collector configuration
	CustomData4 *string `json:"custom_data_4,omitempty"`
	// Free form data set via the customData5 field in the analytics collector configuration
	CustomData5 *string `json:"custom_data_5,omitempty"`
	// Free form data set via the customData6 field in the analytics collector configuration
	CustomData6 *string `json:"custom_data_6,omitempty"`
	// Free form data set via the customData7 field in the analytics collector configuration
	CustomData7 *string `json:"custom_data_7,omitempty"`
	// Free form data set via the customData8 field in the analytics collector configuration
	CustomData8 *string `json:"custom_data_8,omitempty"`
	// Free form data set via the customData9 field in the analytics collector configuration
	CustomData9 *string `json:"custom_data_9,omitempty"`
	// Free form data set via the customData10 field in the analytics collector configuration
	CustomData10 *string `json:"custom_data_10,omitempty"`
	// Free form data set via the customData11 field in the analytics collector configuration
	CustomData11 *string `json:"custom_data_11,omitempty"`
	// Free form data set via the customData12 field in the analytics collector configuration
	CustomData12 *string `json:"custom_data_12,omitempty"`
	// Free form data set via the customData13 field in the analytics collector configuration
	CustomData13 *string `json:"custom_data_13,omitempty"`
	// Free form data set via the customData14 field in the analytics collector configuration
	CustomData14 *string `json:"custom_data_14,omitempty"`
	// Free form data set via the customData15 field in the analytics collector configuration
	CustomData15 *string `json:"custom_data_15,omitempty"`
	// Free form data set via the customData16 field in the analytics collector configuration
	CustomData16 *string `json:"custom_data_16,omitempty"`
	// Free form data set via the customData17 field in the analytics collector configuration
	CustomData17 *string `json:"custom_data_17,omitempty"`
	// Free form data set via the customData18 field in the analytics collector configuration
	CustomData18 *string `json:"custom_data_18,omitempty"`
	// Free form data set via the customData19 field in the analytics collector configuration
	CustomData19 *string `json:"custom_data_19,omitempty"`
	// Free form data set via the customData20 field in the analytics collector configuration
	CustomData20 *string `json:"custom_data_20,omitempty"`
	// Free form data set via the customData21 field in the analytics collector configuration
	CustomData21 *string `json:"custom_data_21,omitempty"`
	// Free form data set via the customData22 field in the analytics collector configuration
	CustomData22 *string `json:"custom_data_22,omitempty"`
	// Free form data set via the customData23 field in the analytics collector configuration
	CustomData23 *string `json:"custom_data_23,omitempty"`
	// Free form data set via the customData24 field in the analytics collector configuration
	CustomData24 *string `json:"custom_data_24,omitempty"`
	// Free form data set via the customData25 field in the analytics collector configuration
	CustomData25 *string `json:"custom_data_25,omitempty"`
	// Free form data set via the customData26 field in the analytics collector configuration
	CustomData26 *string `json:"custom_data_26,omitempty"`
	// Free form data set via the customData27 field in the analytics collector configuration
	CustomData27 *string `json:"custom_data_27,omitempty"`
	// Free form data set via the customData28 field in the analytics collector configuration
	CustomData28 *string `json:"custom_data_28,omitempty"`
	// Free form data set via the customData29 field in the analytics collector configuration
	CustomData29 *string `json:"custom_data_29,omitempty"`
	// Free form data set via the customData30 field in the analytics collector configuration
	CustomData30 *string `json:"custom_data_30,omitempty"`
	// Type of device (Desktop, Phone, Tablet)
	DeviceClass *string `json:"device_class,omitempty"`
	// Type of the device detected via User Agent
	DeviceType *string `json:"device_type,omitempty"`
	// Domain the player was loaded on (.www is stripped away)
	Domain *string `json:"domain,omitempty"`
	// Time in milliseconds it took the DRM server to respond
	DrmLoadTime *int64 `json:"drm_load_time,omitempty"`
	// DRM system used for this impression
	DrmType *string `json:"drm_type,omitempty"`
	// Dropped frames during playback
	DroppedFrames *int64 `json:"dropped_frames,omitempty"`
	// Duration of the sample in milliseconds
	Duration *int64 `json:"duration,omitempty"`
	// Error code
	ErrorCode *int32 `json:"error_code,omitempty"`
	// Error message
	ErrorMessage *string `json:"error_message,omitempty"`
	// A/B test experiment name
	ExperimentName *string `json:"experiment_name,omitempty"`
	// Random UUID that is used to identify a sessions (required)
	ImpressionId *string `json:"impression_id,omitempty"`
	// IP Address of the client
	IpAddress *string `json:"ip_address,omitempty"`
	// Is chromecast active
	IsCasting *bool `json:"is_casting,omitempty"`
	// Is the stream live or VoD
	IsLive *bool `json:"is_live,omitempty"`
	// Is the player muted
	IsMuted *bool `json:"is_muted,omitempty"`
	// The users Internet Service Provider inferred via the IP address
	Isp *string `json:"isp,omitempty"`
	// Language set in the browser
	Language *string `json:"language,omitempty"`
	// Analytics license key
	LicenseKey *string `json:"license_key,omitempty"`
	// URL of the HLS source
	M3u8Url *string `json:"m3u8_url,omitempty"`
	// URL of the DASH source
	MpdUrl *string `json:"mpd_url,omitempty"`
	// Operating system
	Operatingsystem *string `json:"operatingsystem,omitempty"`
	// Operating system version major
	OperatingsystemVersionMajor *string `json:"operatingsystem_version_major,omitempty"`
	// Operating system version minor
	OperatingsystemVersionMinor *string `json:"operatingsystem_version_minor,omitempty"`
	// Time in milliseconds the page took to load
	PageLoadTime *int32 `json:"page_load_time,omitempty"`
	// Player load type. 1 = Foreground, 2 = Background
	PageLoadType *int32 `json:"page_load_type,omitempty"`
	// path on the website
	Path *string `json:"path,omitempty"`
	// Milliseconds the player was paused
	Paused *int64 `json:"paused,omitempty"`
	// Platform the player is running on (web, android, ios)
	Platform *string `json:"platform,omitempty"`
	// Milliseconds the player played
	Played *int64 `json:"played,omitempty"`
	// Video player being used for this session
	Player *string `json:"player,omitempty"`
	// Player license key
	PlayerKey *string `json:"player_key,omitempty"`
	// Time in milliseconds the player took to start up
	PlayerStartuptime *int32 `json:"player_startuptime,omitempty"`
	// HTML or native playback
	PlayerTech *string `json:"player_tech,omitempty"`
	// Player software version
	PlayerVersion *string `json:"player_version,omitempty"`
	// URL of the progressive MP4 source
	ProgUrl *string `json:"prog_url,omitempty"`
	// Geographic region (ISO 3166-2 code)
	Region *string `json:"region,omitempty"`
	// Screen as reported by the browser
	ScreenHeight *int32 `json:"screen_height,omitempty"`
	// Screen as reported by the browser
	ScreenWidth *int32 `json:"screen_width,omitempty"`
	// Milliseconds it took the player to seek
	Seeked *int64 `json:"seeked,omitempty"`
	// Number of video segments downloaded
	SegmentDownloadCount *int32 `json:"segment_download_count,omitempty"`
	// Size of downloaded video segments (byte)
	SegmentDownloadSize *int32 `json:"segment_download_size,omitempty"`
	// Cumulative time needed to download video segments
	SegmentDownloadTime *int32 `json:"segment_download_time,omitempty"`
	// Sequence number of the sample in which it occured in the session
	SequenceNumber *int64 `json:"sequence_number,omitempty"`
	// Video size (FULLSCREEN or WINDOW)
	Size *string `json:"size,omitempty"`
	// Combination of player- and videoStartuptime
	Startuptime *int32 `json:"startuptime,omitempty"`
	// Internal state of the analytics state machine
	State *string `json:"state,omitempty"`
	// Format of the stream (HLS, DASH, Progressive MP4)
	StreamFormat *string `json:"stream_format,omitempty"`
	// Subtitle enabled
	SubtitleEnabled *bool `json:"subtitle_enabled,omitempty"`
	// Selected subtitle language
	SubtitleLanguage *string `json:"subtitle_language,omitempty"`
	// Video codecs supported by platform/browser
	SupportedVideoCodes []string `json:"supported_video_codes,omitempty"`
	// Current time in milliseconds
	Time *int64 `json:"time,omitempty"`
	// ID that is persisted across sessions to identify a browser
	UserId *string `json:"user_id,omitempty"`
	// Bitrate of the played back video rendition
	VideoBitrate *int64 `json:"video_bitrate,omitempty"`
	// Video codec of current stream
	VideoCodec *string `json:"video_codec,omitempty"`
	// Length of the video in milliseconds
	VideoDuration *int64 `json:"video_duration,omitempty"`
	// ID of the video as configured via the analytics config
	VideoId *string `json:"video_id,omitempty"`
	// Free form human readable video title as configured in the analytics config
	VideoTitle *string `json:"video_title,omitempty"`
	// Resolution of the played back Video Rendition
	VideoPlaybackHeight *int32 `json:"video_playback_height,omitempty"`
	// Resolution of the played back Video Rendition
	VideoPlaybackWidth *int32 `json:"video_playback_width,omitempty"`
	// Time in milliseconds it took to start video playback
	VideoStartuptime *int64 `json:"video_startuptime,omitempty"`
	// End time of the sample in the video (milliseconds)
	VideotimeEnd *int64 `json:"videotime_end,omitempty"`
	// Start time of the sample in the video (milliseconds)
	VideotimeStart *int64 `json:"videotime_start,omitempty"`
	// Height of the video player on the page
	VideoWindowHeight *int32 `json:"video_window_height,omitempty"`
	// Width of the video player on the page
	VideoWindowWidth *int32 `json:"video_window_width,omitempty"`
	// True if starting the video failed
	VideostartFailed *bool `json:"videostart_failed,omitempty"`
	// Reason why starting the video failed
	VideostartFailedReason AnalyticsVideoStartFailedReason `json:"videostart_failed_reason,omitempty"`
}
