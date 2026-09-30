package model

// LevelAv1 : Specified set of constraints that indicate a degree of required decoder performance for a profile, see: https://aomediacodec.github.io/av1-spec/av1-spec.pdf (Annex A.3)
type LevelAv1 string

// List of possible LevelAv1 values
const (
	LevelAv1_L2_0 LevelAv1 = "2.0"
	LevelAv1_L2_1 LevelAv1 = "2.1"
	LevelAv1_L3_0 LevelAv1 = "3.0"
	LevelAv1_L3_1 LevelAv1 = "3.1"
	LevelAv1_L4_0 LevelAv1 = "4.0"
	LevelAv1_L4_1 LevelAv1 = "4.1"
	LevelAv1_L5_0 LevelAv1 = "5.0"
	LevelAv1_L5_1 LevelAv1 = "5.1"
	LevelAv1_L5_2 LevelAv1 = "5.2"
	LevelAv1_L5_3 LevelAv1 = "5.3"
	LevelAv1_L6_0 LevelAv1 = "6.0"
	LevelAv1_L6_1 LevelAv1 = "6.1"
	LevelAv1_L6_2 LevelAv1 = "6.2"
	LevelAv1_L6_3 LevelAv1 = "6.3"
)
