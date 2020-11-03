package model


// LimitReferences : LimitReferences model
type LimitReferences string

// List of possible LimitReferences values
const (
    LimitReferences_DISABLED LimitReferences = "DISABLED"
    LimitReferences_DEPTH LimitReferences = "DEPTH"
    LimitReferences_CU LimitReferences = "CU"
    LimitReferences_DEPTH_AND_CU LimitReferences = "DEPTH_AND_CU"
)


