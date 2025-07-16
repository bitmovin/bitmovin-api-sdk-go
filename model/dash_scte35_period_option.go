package model

// DashScte35PeriodOption : Controls how SCTE-35 signals are applied at the DASH period level. <ul>   <li>SINGLE_PERIOD: Applies SCTE markers to a single DASH period.</li>   <li>MULTI_PERIOD: Multiple periods are created based on received SCTE-35 messages.</li> </ul>
type DashScte35PeriodOption string

// List of possible DashScte35PeriodOption values
const (
	DashScte35PeriodOption_SINGLE_PERIOD DashScte35PeriodOption = "SINGLE_PERIOD"
	DashScte35PeriodOption_MULTI_PERIOD  DashScte35PeriodOption = "MULTI_PERIOD"
)
