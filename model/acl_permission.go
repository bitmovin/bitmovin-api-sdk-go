package model

// AclPermission : AclPermission model
type AclPermission string

// List of possible AclPermission values
const (
	AclPermission_PUBLIC_READ AclPermission = "PUBLIC_READ"
	AclPermission_PRIVATE     AclPermission = "PRIVATE"
)
