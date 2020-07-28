package model
// DolbyAtmosMeteringMode : Algorithm to be used for measuring loudness
type DolbyAtmosMeteringMode string

// List of DolbyAtmosMeteringMode
const (
	DolbyAtmosMeteringMode_ITU_R_BS_1770_1 DolbyAtmosMeteringMode = "ITU-R BS.1770-1"
	DolbyAtmosMeteringMode_ITU_R_BS_1770_2 DolbyAtmosMeteringMode = "ITU-R BS.1770-2"
	DolbyAtmosMeteringMode_ITU_R_BS_1770_3 DolbyAtmosMeteringMode = "ITU-R BS.1770-3"
	DolbyAtmosMeteringMode_ITU_R_BS_1770_4 DolbyAtmosMeteringMode = "ITU-R BS.1770-4"
	DolbyAtmosMeteringMode_LEQ_A DolbyAtmosMeteringMode = "Leq (A)"
)