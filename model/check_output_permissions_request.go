package model

// CheckOutputPermissionsRequest model
type CheckOutputPermissionsRequest struct {
	// The path on the storage for which permissions should be checked. In AWS S3 terminology, this corresponds to a \"prefix\". To perform the check, an empty test file (WritePermissionTestFile.txt) will be created in this location. Defaults to an empty string, which corresponds to the root directory.
	Path *string `json:"path,omitempty"`
}
