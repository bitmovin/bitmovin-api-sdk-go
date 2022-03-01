package model

// AnalyticsAdsImpressionSample model
type AnalyticsAdsImpressionSample struct {
	// Ad click through url
	AdClickthroughUrl *string `json:"ad_clickthrough_url,omitempty"`
	// Ad description
	AdDescription *string `json:"ad_description,omitempty"`
	// Ad duration
	AdDuration *int64 `json:"ad_duration,omitempty"`
	// Ad fallback index
	AdFallbackIndex *int64 `json:"ad_fallback_index,omitempty"`
	// Ad id
	AdId *string `json:"ad_id,omitempty"`
	// Ad id player
	AdIdPlayer *string `json:"ad_id_player,omitempty"`
	// Ad impression id
	AdImpressionId *string `json:"ad_impression_id,omitempty"`
	// Ad is persistent
	AdIsPersistent *bool `json:"ad_is_persistent,omitempty"`
	// Ad module
	AdModule *string `json:"ad_module,omitempty"`
	// Ad module version
	AdModuleVersion *string `json:"ad_module_version,omitempty"`
	// Ad offset
	AdOffset *string `json:"ad_offset,omitempty"`
	// Ad playback height
	AdPlaybackHeight *int64 `json:"ad_playback_height,omitempty"`
	// Ad playback width
	AdPlaybackWidth *int64 `json:"ad_playback_width,omitempty"`
	// Ad pod position
	AdPodPosition *int64 `json:"ad_pod_position,omitempty"`
	// Ad position
	AdPosition *string `json:"ad_position,omitempty"`
	// Ad preload offset
	AdPreloadOffset *int64 `json:"ad_preload_offset,omitempty"`
	// Ad replace content duration
	AdReplaceContentDuration *int64 `json:"ad_replace_content_duration,omitempty"`
	// Ad schedule duration
	AdScheduleTime *int64 `json:"ad_schedule_time,omitempty"`
	// Ad skip after
	AdSkipAfter *int64 `json:"ad_skip_after,omitempty"`
	// Ad is skippable
	AdSkippable *bool `json:"ad_skippable,omitempty"`
	// Ad startup time
	AdStartupTime *int64 `json:"ad_startup_time,omitempty"`
	// Ad system
	AdSystem *string `json:"ad_system,omitempty"`
	// Ad tag path
	AdTagPath *string `json:"ad_tag_path,omitempty"`
	// Ad system
	AdTagServer *string `json:"ad_tag_server,omitempty"`
	// Ad tag type
	AdTagType *string `json:"ad_tag_type,omitempty"`
	// Ad tag url
	AdTagUrl *string `json:"ad_tag_url,omitempty"`
	// Ad title
	AdTitle *string `json:"ad_title,omitempty"`
	// Ad wrapper ads count
	AdWrapperAdsCount *int64 `json:"ad_wrapper_ads_count,omitempty"`
	// Advertiser name
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	// Analytics version
	AnalyticsVersion *string `json:"analytics_version,omitempty"`
	// Api framework
	ApiFramework *string `json:"api_framework,omitempty"`
	// Organization id
	ApiorgId *string `json:"apiorg_id,omitempty"`
	// User id
	ApiuserId *string `json:"apiuser_id,omitempty"`
	// Audio bitrate
	AudioBitrate *int64 `json:"audio_bitrate,omitempty"`
	// Is autoplay
	Autoplay *bool `json:"autoplay,omitempty"`
	// Browser name
	Browser *string `json:"browser,omitempty"`
	// Browser is bot
	BrowserIsBot *bool `json:"browser_is_bot,omitempty"`
	// Browser version major
	BrowserVersionMajor *string `json:"browser_version_major,omitempty"`
	// Browser version minor
	BrowserVersionMinor *string `json:"browser_version_minor,omitempty"`
	// CDN Provider
	CdnProvider *string `json:"cdn_provider,omitempty"`
	// City
	City *string `json:"city,omitempty"`
	// Click percentage
	ClickPercentage *int64 `json:"click_percentage,omitempty"`
	// Click position
	ClickPosition *int64 `json:"click_position,omitempty"`
	// Clicked
	Clicked *int32 `json:"clicked,omitempty"`
	// Current time of the client
	ClientTime *int64 `json:"client_time,omitempty"`
	// Close percentage
	ClosePercentage *int64 `json:"close_percentage,omitempty"`
	// Close position
	ClosePosition *int64 `json:"close_position,omitempty"`
	// Closed
	Closed *int32 `json:"closed,omitempty"`
	// Completed
	Completed *int32 `json:"completed,omitempty"`
	// Country
	Country *string `json:"country,omitempty"`
	// Creative ad id
	CreativeAdId *string `json:"creative_ad_id,omitempty"`
	// Creative id
	CreativeId *string `json:"creative_id,omitempty"`
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
	// Custom user ID
	CustomUserId *string `json:"custom_user_id,omitempty"`
	// Deal id
	DealId *string `json:"deal_id,omitempty"`
	// Type of device (Desktop, Phone, Tablet)
	DeviceClass *string `json:"device_class,omitempty"`
	// Type of the device detected via User Agent
	DeviceType *string `json:"device_type,omitempty"`
	// Domain the player was loaded on (.www is stripped away)
	Domain *string `json:"domain,omitempty"`
	// Error code
	ErrorCode *int32 `json:"error_code,omitempty"`
	// Error data
	ErrorData *string `json:"error_data,omitempty"`
	// Error message
	ErrorMessage *string `json:"error_message,omitempty"`
	// Error percentage
	ErrorPercentage *int32 `json:"error_percentage,omitempty"`
	// Error position
	ErrorPosition *int64 `json:"error_position,omitempty"`
	// Exit position
	ExitPosition *int64 `json:"exit_position,omitempty"`
	// A/B test experiment name
	ExperimentName *string `json:"experiment_name,omitempty"`
	// IP Address of the client
	IpAddress *string `json:"ip_address,omitempty"`
	// The users Internet Service Provider inferred via the IP address
	Isp *string `json:"isp,omitempty"`
	// Language set in the browser
	Language *string `json:"language,omitempty"`
	// Analytics license key
	LicenseKey *string `json:"license_key,omitempty"`
	// Manifest download time
	ManifestDownloadTime *int64 `json:"manifest_download_time,omitempty"`
	// Media path
	MediaPath *string `json:"media_path,omitempty"`
	// Media server
	MediaServer *string `json:"media_server,omitempty"`
	// Media url
	MediaUrl *string `json:"media_url,omitempty"`
	// Midpoint
	Midpoint *int32 `json:"midpoint,omitempty"`
	// Minimum suggested duration
	MinSuggestedDuration *int64 `json:"min_suggested_duration,omitempty"`
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
	// Percentage in viewport
	PercentageInViewport *int64 `json:"percentage_in_viewport,omitempty"`
	// Platform the player is running on (web, android, ios)
	Platform *string `json:"platform,omitempty"`
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
	// Play percentage
	PlayPercentage *int32 `json:"play_percentage,omitempty"`
	// Quartile 1
	Quartile1 *int32 `json:"quartile_1,omitempty"`
	// Quartile 3
	Quartile3 *int32 `json:"quartile_3,omitempty"`
	// Geographic region (ISO 3166-2 code)
	Region *string `json:"region,omitempty"`
	// Screen as reported by the browser
	ScreenHeight *int32 `json:"screen_height,omitempty"`
	// Screen as reported by the browser
	ScreenWidth *int32 `json:"screen_width,omitempty"`
	// Video size (FULLSCREEN or WINDOW)
	Size *string `json:"size,omitempty"`
	// Skip percentage
	SkipPercentage *int32 `json:"skip_percentage,omitempty"`
	// Skip position
	SkipPosition *int64 `json:"skip_position,omitempty"`
	// Skipped
	Skipped *int32 `json:"skipped,omitempty"`
	// Started
	Started *int32 `json:"started,omitempty"`
	// Format of the stream (HLS, DASH, Progressive MP4)
	StreamFormat *string `json:"stream_format,omitempty"`
	// Survey url
	SurveyUrl *string `json:"survey_url,omitempty"`
	// Current time in milliseconds
	Time *int64 `json:"time,omitempty"`
	// Time in viewport
	TimeInViewport *int64 `json:"time_in_viewport,omitempty"`
	// Time played
	TimePlayed *int64 `json:"time_played,omitempty"`
	// Universal ad id registry
	UniversalAdIdRegistry *string `json:"universal_ad_id_registry,omitempty"`
	// Universal ad id value
	UniversalAdIdValue *string `json:"universal_ad_id_value,omitempty"`
	// ID that is persisted across sessions to identify a browser
	UserId *string `json:"user_id,omitempty"`
	// Bitrate of the played back video rendition
	VideoBitrate *int64 `json:"video_bitrate,omitempty"`
	// ID of the video
	VideoId *string `json:"video_id,omitempty"`
	// ID of related video impression
	VideoImpressionId *string `json:"video_impression_id,omitempty"`
	// Free form human readable video ad title
	VideoTitle *string `json:"video_title,omitempty"`
	// Height of the video player on the page
	VideoWindowHeight *int32 `json:"video_window_height,omitempty"`
	// Width of the video player on the page
	VideoWindowWidth *int32 `json:"video_window_width,omitempty"`
}
