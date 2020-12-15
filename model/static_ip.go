package model

// StaticIp model
type StaticIp struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// The IPv4 address of the static ip
	IpAddress *string `json:"ipAddress,omitempty"`
	// Required if the static IP should be created for an AWS infrastructure account.
	InfrastructureId *string `json:"infrastructureId,omitempty"`
	// Status of the Static Ip
	Status StaticIpStatus `json:"status,omitempty"`
	// The region of the static Ip (required)
	Region CloudRegion `json:"region,omitempty"`
}
