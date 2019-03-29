package model

type PlayerThirdPartyLicensing struct {
	// URL to your license check server
	LicenseCheckServer string `json:"licenseCheckServer,omitempty"`
	// Timeout in ms
	LicenseCheckTimeout *int32 `json:"licenseCheckTimeout,omitempty"`
	// Specify if the Licensing Request should fail or not on Third Party Licensing Error
	ErrorAction PlayerThirdPartyLicensingErrorAction `json:"errorAction,omitempty"`
	// Specify if the Licensing Request should fail or not on Third Party Licensing timeout
	TimeoutAction PlayerThirdPartyLicensingErrorAction `json:"timeoutAction,omitempty"`
}

