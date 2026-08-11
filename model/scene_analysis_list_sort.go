package model

// SceneAnalysisListSort : SceneAnalysisListSort model
type SceneAnalysisListSort string

// List of possible SceneAnalysisListSort values
const (
	SceneAnalysisListSort_CREATED_AT_DESC SceneAnalysisListSort = "createdAt:DESC"
	SceneAnalysisListSort_CREATED_AT_ASC  SceneAnalysisListSort = "createdAt:ASC"
	SceneAnalysisListSort_RELEVANCE_DESC  SceneAnalysisListSort = "relevance:DESC"
)
