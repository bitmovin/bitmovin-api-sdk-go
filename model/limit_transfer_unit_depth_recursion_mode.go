package model
// LimitTransferUnitDepthRecursionMode : Enables early exit from transfer unit depth recursion, for inter coded blocks.
type LimitTransferUnitDepthRecursionMode string

// List of LimitTransferUnitDepthRecursionMode
const (
	LimitTransferUnitDepthRecursionMode_DISABLED LimitTransferUnitDepthRecursionMode = "DISABLED"
	LimitTransferUnitDepthRecursionMode_LEVEL_1 LimitTransferUnitDepthRecursionMode = "LEVEL_1"
	LimitTransferUnitDepthRecursionMode_LEVEL_2 LimitTransferUnitDepthRecursionMode = "LEVEL_2"
	LimitTransferUnitDepthRecursionMode_LEVEL_3 LimitTransferUnitDepthRecursionMode = "LEVEL_3"
	LimitTransferUnitDepthRecursionMode_LEVEL_4 LimitTransferUnitDepthRecursionMode = "LEVEL_4"
)