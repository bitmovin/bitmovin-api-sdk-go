package model

// SceneType : SceneType model
type SceneType string

// List of possible SceneType values
const (
	SceneType_LOGO_IDENT            SceneType = "LOGO_IDENT"
	SceneType_OPENING_CREDITS       SceneType = "OPENING_CREDITS"
	SceneType_RECAP                 SceneType = "RECAP"
	SceneType_PREVIEW_THIS_TITLE    SceneType = "PREVIEW_THIS_TITLE"
	SceneType_PROMOTION_OTHER_TITLE SceneType = "PROMOTION_OTHER_TITLE"
	SceneType_BREAK_BUMPER          SceneType = "BREAK_BUMPER"
	SceneType_END_CREDITS           SceneType = "END_CREDITS"
	SceneType_ADS                   SceneType = "ADS"
	SceneType_MAIN_CONTENT          SceneType = "MAIN_CONTENT"
	SceneType_UNKNOWN               SceneType = "UNKNOWN"
)
