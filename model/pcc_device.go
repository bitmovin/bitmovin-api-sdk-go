package model

// PccDevice model
type PccDevice struct {
	// The device pool, as the reader is shown it. (required)
	Name *string `json:"name,omitempty"`
	// True where no naming rule recognised this pool, so `name` is a stated placeholder rather than the pool's own. The units and sessions below still tell two such pools apart. (required)
	NamePlaceholder *bool `json:"namePlaceholder,omitempty"`
	// What distinguishes this pool from another of the same name, where anything does.
	Qualifier *string `json:"qualifier,omitempty"`
	// The fleet's own classification, such as `tv`, `desktop`, `stb` or `mobile`. Never one guessed from a name, and not a closed set: the fleet may answer with a kind this list does not name.
	DeviceType *string         `json:"deviceType,omitempty"`
	SessionIds []string        `json:"sessionIds,omitempty"`
	Units      []PccDeviceUnit `json:"units,omitempty"`
	// Every player version that measured this pool. More than one means the runs spanned a player release. (required)
	PlayerVersions []string `json:"playerVersions,omitempty"`
	// One per column, in `combinations` order. (required)
	Cells []PccCell `json:"cells,omitempty"`
}
