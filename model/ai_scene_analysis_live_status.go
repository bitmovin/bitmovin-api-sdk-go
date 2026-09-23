package model

// AiSceneAnalysisLiveStatus : AiSceneAnalysisLiveStatus model
type AiSceneAnalysisLiveStatus string

// List of possible AiSceneAnalysisLiveStatus values
const (
	AiSceneAnalysisLiveStatus_CREATED        AiSceneAnalysisLiveStatus = "CREATED"
	AiSceneAnalysisLiveStatus_QUEUED         AiSceneAnalysisLiveStatus = "QUEUED"
	AiSceneAnalysisLiveStatus_RUNNING        AiSceneAnalysisLiveStatus = "RUNNING"
	AiSceneAnalysisLiveStatus_FINISHED       AiSceneAnalysisLiveStatus = "FINISHED"
	AiSceneAnalysisLiveStatus_CANCELED       AiSceneAnalysisLiveStatus = "CANCELED"
	AiSceneAnalysisLiveStatus_ERROR          AiSceneAnalysisLiveStatus = "ERROR"
	AiSceneAnalysisLiveStatus_TRANSFER_ERROR AiSceneAnalysisLiveStatus = "TRANSFER_ERROR"
)
