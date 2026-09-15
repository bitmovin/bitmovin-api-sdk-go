package model

// PccDeviceUnit model
type PccDeviceUnit struct {
	// The fleet's identifier for one physical machine, so two rows of the same pool can be told apart. (required)
	UnitId         *string  `json:"unitId,omitempty"`
	SessionIds     []string `json:"sessionIds,omitempty"`
	BrowserVersion *string  `json:"browserVersion,omitempty"`
	OsVersion      *string  `json:"osVersion,omitempty"`
	// Attribute tags such as `webos:firmwareVersion:33.23.05`, which is where a television's firmware lives.
	Tags []string `json:"tags,omitempty"`
}
