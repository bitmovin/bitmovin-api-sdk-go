package model

// SceneType : SceneType model
type SceneType string

// List of possible SceneType values
const (
	SceneType_OPENING_TITLES            SceneType = "OPENING_TITLES"
	SceneType_EPISODE_TITLE_CARD        SceneType = "EPISODE_TITLE_CARD"
	SceneType_STUDIO_LOGO_BUMPER        SceneType = "STUDIO_LOGO_BUMPER"
	SceneType_NETWORK_OR_PLATFORM_IDENT SceneType = "NETWORK_OR_PLATFORM_IDENT"
	SceneType_RECAP                     SceneType = "RECAP"
	SceneType_PREVIEW_THIS_TITLE        SceneType = "PREVIEW_THIS_TITLE"
	SceneType_PROMO_OTHER_TITLE         SceneType = "PROMO_OTHER_TITLE"
	SceneType_TRAILER_OTHER_TITLE       SceneType = "TRAILER_OTHER_TITLE"
	SceneType_ADS                       SceneType = "ADS"
	SceneType_ACT_BREAK_EYECATCH        SceneType = "ACT_BREAK_EYECATCH"
	SceneType_TECHNICAL_SLATE_OR_TEST   SceneType = "TECHNICAL_SLATE_OR_TEST"
	SceneType_MAIN_CONTENT              SceneType = "MAIN_CONTENT"
	SceneType_MID_CREDIT_SCENE          SceneType = "MID_CREDIT_SCENE"
	SceneType_POST_CREDIT_SCENE         SceneType = "POST_CREDIT_SCENE"
	SceneType_END_CREDITS               SceneType = "END_CREDITS"
	SceneType_UNKNOWN                   SceneType = "UNKNOWN"
)
