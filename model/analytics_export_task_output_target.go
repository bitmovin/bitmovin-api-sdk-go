package model

type AnalyticsExportTaskOutputTarget struct {
	// Path where the export should be saved
	OutputPath string `json:"outputPath,omitempty"`
	// Id of the output that should be used
	OutputId string `json:"outputId,omitempty"`
}

