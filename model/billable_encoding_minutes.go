package model

type BillableEncodingMinutes struct {
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	Codec CodecConfigType `json:"codec,omitempty"`
	PerTitleResultStream StatisticsPerTitleStream `json:"perTitleResultStream,omitempty"`
	PsnrMode PsnrPerStreamMode `json:"psnrMode,omitempty"`
	// Details about billable minutes for each resolution category
	BillableMinutesDetails []BillableEncodingMinutesDetails `json:"billableMinutesDetails,omitempty"`
}

