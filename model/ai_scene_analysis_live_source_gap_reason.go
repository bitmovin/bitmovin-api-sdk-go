package model

// AiSceneAnalysisLiveSourceGapReason : AiSceneAnalysisLiveSourceGapReason model
type AiSceneAnalysisLiveSourceGapReason string

// List of possible AiSceneAnalysisLiveSourceGapReason values
const (
	AiSceneAnalysisLiveSourceGapReason_SOURCE_DISCONNECTED       AiSceneAnalysisLiveSourceGapReason = "SOURCE_DISCONNECTED"
	AiSceneAnalysisLiveSourceGapReason_PROCESSING_MEDIA_PRESSURE AiSceneAnalysisLiveSourceGapReason = "PROCESSING_MEDIA_PRESSURE"
	AiSceneAnalysisLiveSourceGapReason_ANALYSIS_LAG              AiSceneAnalysisLiveSourceGapReason = "ANALYSIS_LAG"
	AiSceneAnalysisLiveSourceGapReason_WINDOW_BUILD_FAILED       AiSceneAnalysisLiveSourceGapReason = "WINDOW_BUILD_FAILED"
	AiSceneAnalysisLiveSourceGapReason_FINALIZATION_BACKLOG      AiSceneAnalysisLiveSourceGapReason = "FINALIZATION_BACKLOG"
)
