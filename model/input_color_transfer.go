package model

// InputColorTransfer : InputColorTransfer model
type InputColorTransfer string

// List of possible InputColorTransfer values
const (
	InputColorTransfer_UNSPECIFIED  InputColorTransfer = "UNSPECIFIED"
	InputColorTransfer_BT709        InputColorTransfer = "BT709"
	InputColorTransfer_GAMMA22      InputColorTransfer = "GAMMA22"
	InputColorTransfer_GAMMA28      InputColorTransfer = "GAMMA28"
	InputColorTransfer_SMPTE170M    InputColorTransfer = "SMPTE170M"
	InputColorTransfer_SMPTE240M    InputColorTransfer = "SMPTE240M"
	InputColorTransfer_LINEAR       InputColorTransfer = "LINEAR"
	InputColorTransfer_LOG          InputColorTransfer = "LOG"
	InputColorTransfer_LOG_SQRT     InputColorTransfer = "LOG_SQRT"
	InputColorTransfer_IEC61966_2_4 InputColorTransfer = "IEC61966_2_4"
	InputColorTransfer_BT1361_ECG   InputColorTransfer = "BT1361_ECG"
	InputColorTransfer_IEC61966_2_1 InputColorTransfer = "IEC61966_2_1"
	InputColorTransfer_BT2020_10    InputColorTransfer = "BT2020_10"
	InputColorTransfer_BT2020_12    InputColorTransfer = "BT2020_12"
	InputColorTransfer_SMPTE2084    InputColorTransfer = "SMPTE2084"
	InputColorTransfer_SMPTE428     InputColorTransfer = "SMPTE428"
	InputColorTransfer_ARIB_STD_B67 InputColorTransfer = "ARIB_STD_B67"
)
