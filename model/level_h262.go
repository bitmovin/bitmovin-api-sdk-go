package model
// LevelH262 : Specified set of constraints that indicate a degree of required decoder performance for a profile, see: https://en.wikipedia.org/wiki/H.264/MPEG-4_AVC#Levels
type LevelH262 string

// List of LevelH262
const (
	LevelH262_MAIN LevelH262 = "MAIN"
	LevelH262_HIGH LevelH262 = "HIGH"
)