package model

// PccCell model
type PccCell struct {
	// What a combination says once every session that measured it has been read. Five of the nine answer for the measurement rather than for the device; `aboutTheDevice` says which, and folding those into \"not supported\" is how this data gets misread. (required)
	Verdict PccVerdict `json:"verdict,omitempty"`
	// The reader's word for that verdict — `Supported`, `Not supported`, `Not measured`, and so on. Fewer words than there are verdicts: three of them read as `Not measured`. `legend` lists every word. (required)
	Label *string `json:"label,omitempty"`
	// False where the verdict says something about the measurement rather than the device. (required)
	AboutTheDevice *bool `json:"aboutTheDevice,omitempty"`
	// The cell's whole account in one paragraph: the verdict, what agreed, the picture, the stream. (required)
	Account *string `json:"account,omitempty"`
	// The grid's own mark for that verdict, which `legend` explains. (required)
	Symbol *string `json:"symbol,omitempty"`
	// `3/4` where a session disagreed with the published verdict, and absent where none did.
	Agreement *string `json:"agreement,omitempty"`
	// What the sessions that disagreed recorded, spelled out. Present only where `agreement` is.
	AgreementAccount *string     `json:"agreementAccount,omitempty"`
	Picture          *PccPicture `json:"picture,omitempty"`
	// The sessions this verdict was taken from. Quote one to Bitmovin support and the measurement behind this cell can be looked up, for as long as the fleet still holds it. (required)
	AgreeingSessionIds []string `json:"agreeingSessionIds,omitempty"`
	// How many sessions recorded anything at all for this combination. (required)
	ResultSessions *float32 `json:"resultSessions,omitempty"`
	// Evidence excluded by the start date or session limit, including pools with no included sessions.
	Recency *string `json:"recency,omitempty"`
}
