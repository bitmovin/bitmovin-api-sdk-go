package model

// Vp9ArnrType : altref noise reduction filter type
type Vp9ArnrType string

// List of possible Vp9ArnrType values
const (
	Vp9ArnrType_BACKWARD Vp9ArnrType = "BACKWARD"
	Vp9ArnrType_FORWARD  Vp9ArnrType = "FORWARD"
	Vp9ArnrType_CENTERED Vp9ArnrType = "CENTERED"
)
