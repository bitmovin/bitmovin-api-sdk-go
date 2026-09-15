package model

// PccSummary model
type PccSummary struct {
	Overview   *PccOverview    `json:"overview,omitempty"`
	CodecReach []PccCodecReach `json:"codecReach,omitempty"`
	// Selected applicable codec/protection combinations with no device-answering verdict. Declared unsupported and claimed-but-not-played are answers; inapplicable pairings are not gaps. (required)
	UnansweredCombinations *int32               `json:"unansweredCombinations,omitempty"`
	Verdicts               []PccVerdictShare    `json:"verdicts,omitempty"`
	ByCodec                []PccSupportShare    `json:"byCodec,omitempty"`
	ByProtection           []PccSupportShare    `json:"byProtection,omitempty"`
	ByDeviceType           []PccDeviceTypeShare `json:"byDeviceType,omitempty"`
	// Every selected combination, with what the selected device pools answered about it. One that several pools claimed and none played points at the stream rather than at the devices. (required)
	Combinations []PccCombinationEvidence `json:"combinations,omitempty"`
	Hdr          *PccHdrSummary           `json:"hdr,omitempty"`
}
