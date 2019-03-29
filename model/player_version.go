package model
import (
	"time"
)

type PlayerVersion struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Version of the Player
	Version string `json:"version,omitempty"`
	// URL of the specified player
	CdnUrl string `json:"cdnUrl,omitempty"`
	// Download URL of the specified player package
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// Creation timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

