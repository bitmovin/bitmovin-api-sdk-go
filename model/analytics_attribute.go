package model

// AnalyticsAttribute : AnalyticsAttribute model
type AnalyticsAttribute string

// List of possible AnalyticsAttribute values
const (
	AnalyticsAttribute_IMPRESSION_ID                  AnalyticsAttribute = "IMPRESSION_ID"
	AnalyticsAttribute_ACTIVE_PLAYER_STARTUPTIME      AnalyticsAttribute = "ACTIVE_PLAYER_STARTUPTIME"
	AnalyticsAttribute_AD                             AnalyticsAttribute = "AD"
	AnalyticsAttribute_ANALYTICS_VERSION              AnalyticsAttribute = "ANALYTICS_VERSION"
	AnalyticsAttribute_AUDIO_BITRATE                  AnalyticsAttribute = "AUDIO_BITRATE"
	AnalyticsAttribute_AUDIO_CODEC                    AnalyticsAttribute = "AUDIO_CODEC"
	AnalyticsAttribute_AUTOPLAY                       AnalyticsAttribute = "AUTOPLAY"
	AnalyticsAttribute_BROWSER                        AnalyticsAttribute = "BROWSER"
	AnalyticsAttribute_BROWSER_VERSION_MAJOR          AnalyticsAttribute = "BROWSER_VERSION_MAJOR"
	AnalyticsAttribute_BROWSER_IS_BOT                 AnalyticsAttribute = "BROWSER_IS_BOT"
	AnalyticsAttribute_BUFFERED                       AnalyticsAttribute = "BUFFERED"
	AnalyticsAttribute_CDN_PROVIDER                   AnalyticsAttribute = "CDN_PROVIDER"
	AnalyticsAttribute_CAST_TECH                      AnalyticsAttribute = "CAST_TECH"
	AnalyticsAttribute_CITY                           AnalyticsAttribute = "CITY"
	AnalyticsAttribute_CLIENT_TIME                    AnalyticsAttribute = "CLIENT_TIME"
	AnalyticsAttribute_COUNTRY                        AnalyticsAttribute = "COUNTRY"
	AnalyticsAttribute_CUSTOM_DATA_1                  AnalyticsAttribute = "CUSTOM_DATA_1"
	AnalyticsAttribute_CUSTOM_DATA_2                  AnalyticsAttribute = "CUSTOM_DATA_2"
	AnalyticsAttribute_CUSTOM_DATA_3                  AnalyticsAttribute = "CUSTOM_DATA_3"
	AnalyticsAttribute_CUSTOM_DATA_4                  AnalyticsAttribute = "CUSTOM_DATA_4"
	AnalyticsAttribute_CUSTOM_DATA_5                  AnalyticsAttribute = "CUSTOM_DATA_5"
	AnalyticsAttribute_CUSTOM_DATA_6                  AnalyticsAttribute = "CUSTOM_DATA_6"
	AnalyticsAttribute_CUSTOM_DATA_7                  AnalyticsAttribute = "CUSTOM_DATA_7"
	AnalyticsAttribute_CUSTOM_DATA_8                  AnalyticsAttribute = "CUSTOM_DATA_8"
	AnalyticsAttribute_CUSTOM_DATA_9                  AnalyticsAttribute = "CUSTOM_DATA_9"
	AnalyticsAttribute_CUSTOM_DATA_10                 AnalyticsAttribute = "CUSTOM_DATA_10"
	AnalyticsAttribute_CUSTOM_DATA_11                 AnalyticsAttribute = "CUSTOM_DATA_11"
	AnalyticsAttribute_CUSTOM_DATA_12                 AnalyticsAttribute = "CUSTOM_DATA_12"
	AnalyticsAttribute_CUSTOM_DATA_13                 AnalyticsAttribute = "CUSTOM_DATA_13"
	AnalyticsAttribute_CUSTOM_DATA_14                 AnalyticsAttribute = "CUSTOM_DATA_14"
	AnalyticsAttribute_CUSTOM_DATA_15                 AnalyticsAttribute = "CUSTOM_DATA_15"
	AnalyticsAttribute_CUSTOM_DATA_16                 AnalyticsAttribute = "CUSTOM_DATA_16"
	AnalyticsAttribute_CUSTOM_DATA_17                 AnalyticsAttribute = "CUSTOM_DATA_17"
	AnalyticsAttribute_CUSTOM_DATA_18                 AnalyticsAttribute = "CUSTOM_DATA_18"
	AnalyticsAttribute_CUSTOM_DATA_19                 AnalyticsAttribute = "CUSTOM_DATA_19"
	AnalyticsAttribute_CUSTOM_DATA_20                 AnalyticsAttribute = "CUSTOM_DATA_20"
	AnalyticsAttribute_CUSTOM_DATA_21                 AnalyticsAttribute = "CUSTOM_DATA_21"
	AnalyticsAttribute_CUSTOM_DATA_22                 AnalyticsAttribute = "CUSTOM_DATA_22"
	AnalyticsAttribute_CUSTOM_DATA_23                 AnalyticsAttribute = "CUSTOM_DATA_23"
	AnalyticsAttribute_CUSTOM_DATA_24                 AnalyticsAttribute = "CUSTOM_DATA_24"
	AnalyticsAttribute_CUSTOM_DATA_25                 AnalyticsAttribute = "CUSTOM_DATA_25"
	AnalyticsAttribute_CUSTOM_DATA_26                 AnalyticsAttribute = "CUSTOM_DATA_26"
	AnalyticsAttribute_CUSTOM_DATA_27                 AnalyticsAttribute = "CUSTOM_DATA_27"
	AnalyticsAttribute_CUSTOM_DATA_28                 AnalyticsAttribute = "CUSTOM_DATA_28"
	AnalyticsAttribute_CUSTOM_DATA_29                 AnalyticsAttribute = "CUSTOM_DATA_29"
	AnalyticsAttribute_CUSTOM_DATA_30                 AnalyticsAttribute = "CUSTOM_DATA_30"
	AnalyticsAttribute_CUSTOM_DATA_31                 AnalyticsAttribute = "CUSTOM_DATA_31"
	AnalyticsAttribute_CUSTOM_DATA_32                 AnalyticsAttribute = "CUSTOM_DATA_32"
	AnalyticsAttribute_CUSTOM_DATA_33                 AnalyticsAttribute = "CUSTOM_DATA_33"
	AnalyticsAttribute_CUSTOM_DATA_34                 AnalyticsAttribute = "CUSTOM_DATA_34"
	AnalyticsAttribute_CUSTOM_DATA_35                 AnalyticsAttribute = "CUSTOM_DATA_35"
	AnalyticsAttribute_CUSTOM_DATA_36                 AnalyticsAttribute = "CUSTOM_DATA_36"
	AnalyticsAttribute_CUSTOM_DATA_37                 AnalyticsAttribute = "CUSTOM_DATA_37"
	AnalyticsAttribute_CUSTOM_DATA_38                 AnalyticsAttribute = "CUSTOM_DATA_38"
	AnalyticsAttribute_CUSTOM_DATA_39                 AnalyticsAttribute = "CUSTOM_DATA_39"
	AnalyticsAttribute_CUSTOM_DATA_40                 AnalyticsAttribute = "CUSTOM_DATA_40"
	AnalyticsAttribute_CUSTOM_DATA_41                 AnalyticsAttribute = "CUSTOM_DATA_41"
	AnalyticsAttribute_CUSTOM_DATA_42                 AnalyticsAttribute = "CUSTOM_DATA_42"
	AnalyticsAttribute_CUSTOM_DATA_43                 AnalyticsAttribute = "CUSTOM_DATA_43"
	AnalyticsAttribute_CUSTOM_DATA_44                 AnalyticsAttribute = "CUSTOM_DATA_44"
	AnalyticsAttribute_CUSTOM_DATA_45                 AnalyticsAttribute = "CUSTOM_DATA_45"
	AnalyticsAttribute_CUSTOM_DATA_46                 AnalyticsAttribute = "CUSTOM_DATA_46"
	AnalyticsAttribute_CUSTOM_DATA_47                 AnalyticsAttribute = "CUSTOM_DATA_47"
	AnalyticsAttribute_CUSTOM_DATA_48                 AnalyticsAttribute = "CUSTOM_DATA_48"
	AnalyticsAttribute_CUSTOM_DATA_49                 AnalyticsAttribute = "CUSTOM_DATA_49"
	AnalyticsAttribute_CUSTOM_DATA_50                 AnalyticsAttribute = "CUSTOM_DATA_50"
	AnalyticsAttribute_CUSTOM_USER_ID                 AnalyticsAttribute = "CUSTOM_USER_ID"
	AnalyticsAttribute_DAY                            AnalyticsAttribute = "DAY"
	AnalyticsAttribute_DEVICE_CLASS                   AnalyticsAttribute = "DEVICE_CLASS"
	AnalyticsAttribute_DEVICE_TYPE                    AnalyticsAttribute = "DEVICE_TYPE"
	AnalyticsAttribute_DOMAIN                         AnalyticsAttribute = "DOMAIN"
	AnalyticsAttribute_DRM_LOAD_TIME                  AnalyticsAttribute = "DRM_LOAD_TIME"
	AnalyticsAttribute_DRM_TYPE                       AnalyticsAttribute = "DRM_TYPE"
	AnalyticsAttribute_DROPPED_FRAMES                 AnalyticsAttribute = "DROPPED_FRAMES"
	AnalyticsAttribute_DURATION                       AnalyticsAttribute = "DURATION"
	AnalyticsAttribute_ERROR_CODE                     AnalyticsAttribute = "ERROR_CODE"
	AnalyticsAttribute_ERROR_MESSAGE                  AnalyticsAttribute = "ERROR_MESSAGE"
	AnalyticsAttribute_ERROR_RATE                     AnalyticsAttribute = "ERROR_RATE"
	AnalyticsAttribute_ERROR_PERCENTAGE               AnalyticsAttribute = "ERROR_PERCENTAGE"
	AnalyticsAttribute_EXPERIMENT_NAME                AnalyticsAttribute = "EXPERIMENT_NAME"
	AnalyticsAttribute_FUNCTION                       AnalyticsAttribute = "FUNCTION"
	AnalyticsAttribute_HOUR                           AnalyticsAttribute = "HOUR"
	AnalyticsAttribute_INITIAL_TIME_TO_TARGET_LATENCY AnalyticsAttribute = "INITIAL_TIME_TO_TARGET_LATENCY"
	AnalyticsAttribute_IP_ADDRESS                     AnalyticsAttribute = "IP_ADDRESS"
	AnalyticsAttribute_IS_CASTING                     AnalyticsAttribute = "IS_CASTING"
	AnalyticsAttribute_IS_LIVE                        AnalyticsAttribute = "IS_LIVE"
	AnalyticsAttribute_IS_LOW_LATENCY                 AnalyticsAttribute = "IS_LOW_LATENCY"
	AnalyticsAttribute_IS_MUTED                       AnalyticsAttribute = "IS_MUTED"
	AnalyticsAttribute_ISP                            AnalyticsAttribute = "ISP"
	AnalyticsAttribute_LANGUAGE                       AnalyticsAttribute = "LANGUAGE"
	AnalyticsAttribute_LATENCY                        AnalyticsAttribute = "LATENCY"
	AnalyticsAttribute_LICENSE_KEY                    AnalyticsAttribute = "LICENSE_KEY"
	AnalyticsAttribute_M3U8_URL                       AnalyticsAttribute = "M3U8_URL"
	AnalyticsAttribute_MINUTE                         AnalyticsAttribute = "MINUTE"
	AnalyticsAttribute_MONTH                          AnalyticsAttribute = "MONTH"
	AnalyticsAttribute_MPD_URL                        AnalyticsAttribute = "MPD_URL"
	AnalyticsAttribute_OPERATINGSYSTEM                AnalyticsAttribute = "OPERATINGSYSTEM"
	AnalyticsAttribute_OPERATINGSYSTEM_VERSION_MAJOR  AnalyticsAttribute = "OPERATINGSYSTEM_VERSION_MAJOR"
	AnalyticsAttribute_PAGE_LOAD_TIME                 AnalyticsAttribute = "PAGE_LOAD_TIME"
	AnalyticsAttribute_PAGE_LOAD_TYPE                 AnalyticsAttribute = "PAGE_LOAD_TYPE"
	AnalyticsAttribute_PATH                           AnalyticsAttribute = "PATH"
	AnalyticsAttribute_PAUSED                         AnalyticsAttribute = "PAUSED"
	AnalyticsAttribute_PLATFORM                       AnalyticsAttribute = "PLATFORM"
	AnalyticsAttribute_PLAY_ATTEMPTS                  AnalyticsAttribute = "PLAY_ATTEMPTS"
	AnalyticsAttribute_PLAYED                         AnalyticsAttribute = "PLAYED"
	AnalyticsAttribute_PLAYER                         AnalyticsAttribute = "PLAYER"
	AnalyticsAttribute_PLAYER_KEY                     AnalyticsAttribute = "PLAYER_KEY"
	AnalyticsAttribute_PLAYER_STARTUPTIME             AnalyticsAttribute = "PLAYER_STARTUPTIME"
	AnalyticsAttribute_PLAYER_TECH                    AnalyticsAttribute = "PLAYER_TECH"
	AnalyticsAttribute_PLAYER_VERSION                 AnalyticsAttribute = "PLAYER_VERSION"
	AnalyticsAttribute_PROG_URL                       AnalyticsAttribute = "PROG_URL"
	AnalyticsAttribute_REBUFFER_PERCENTAGE            AnalyticsAttribute = "REBUFFER_PERCENTAGE"
	AnalyticsAttribute_REGION                         AnalyticsAttribute = "REGION"
	AnalyticsAttribute_SCALE_FACTOR                   AnalyticsAttribute = "SCALE_FACTOR"
	AnalyticsAttribute_SCREEN_HEIGHT                  AnalyticsAttribute = "SCREEN_HEIGHT"
	AnalyticsAttribute_SCREEN_WIDTH                   AnalyticsAttribute = "SCREEN_WIDTH"
	AnalyticsAttribute_SCREEN_ORIENTATION             AnalyticsAttribute = "SCREEN_ORIENTATION"
	AnalyticsAttribute_SEEKED                         AnalyticsAttribute = "SEEKED"
	AnalyticsAttribute_SEQUENCE_NUMBER                AnalyticsAttribute = "SEQUENCE_NUMBER"
	AnalyticsAttribute_SIZE                           AnalyticsAttribute = "SIZE"
	AnalyticsAttribute_STARTUPTIME                    AnalyticsAttribute = "STARTUPTIME"
	AnalyticsAttribute_STREAM_FORMAT                  AnalyticsAttribute = "STREAM_FORMAT"
	AnalyticsAttribute_SUPPORTED_VIDEO_CODECS         AnalyticsAttribute = "SUPPORTED_VIDEO_CODECS"
	AnalyticsAttribute_TARGET_LATENCY                 AnalyticsAttribute = "TARGET_LATENCY"
	AnalyticsAttribute_TARGET_LATENCY_DELTA           AnalyticsAttribute = "TARGET_LATENCY_DELTA"
	AnalyticsAttribute_TIME                           AnalyticsAttribute = "TIME"
	AnalyticsAttribute_TIME_TO_TARGET_LATENCY         AnalyticsAttribute = "TIME_TO_TARGET_LATENCY"
	AnalyticsAttribute_USER_ID                        AnalyticsAttribute = "USER_ID"
	AnalyticsAttribute_VIDEO_BITRATE                  AnalyticsAttribute = "VIDEO_BITRATE"
	AnalyticsAttribute_VIDEO_CODEC                    AnalyticsAttribute = "VIDEO_CODEC"
	AnalyticsAttribute_VIDEO_DURATION                 AnalyticsAttribute = "VIDEO_DURATION"
	AnalyticsAttribute_VIDEO_ID                       AnalyticsAttribute = "VIDEO_ID"
	AnalyticsAttribute_VIDEO_PLAYBACK_HEIGHT          AnalyticsAttribute = "VIDEO_PLAYBACK_HEIGHT"
	AnalyticsAttribute_VIDEO_PLAYBACK_WIDTH           AnalyticsAttribute = "VIDEO_PLAYBACK_WIDTH"
	AnalyticsAttribute_VIDEO_STARTUPTIME              AnalyticsAttribute = "VIDEO_STARTUPTIME"
	AnalyticsAttribute_VIDEO_TITLE                    AnalyticsAttribute = "VIDEO_TITLE"
	AnalyticsAttribute_VIDEO_WINDOW_HEIGHT            AnalyticsAttribute = "VIDEO_WINDOW_HEIGHT"
	AnalyticsAttribute_VIDEO_WINDOW_WIDTH             AnalyticsAttribute = "VIDEO_WINDOW_WIDTH"
	AnalyticsAttribute_VIDEOTIME_END                  AnalyticsAttribute = "VIDEOTIME_END"
	AnalyticsAttribute_VIDEOTIME_START                AnalyticsAttribute = "VIDEOTIME_START"
	AnalyticsAttribute_VIDEOSTART_FAILED              AnalyticsAttribute = "VIDEOSTART_FAILED"
	AnalyticsAttribute_VIDEOSTART_FAILED_REASON       AnalyticsAttribute = "VIDEOSTART_FAILED_REASON"
	AnalyticsAttribute_VIEWTIME                       AnalyticsAttribute = "VIEWTIME"
)
