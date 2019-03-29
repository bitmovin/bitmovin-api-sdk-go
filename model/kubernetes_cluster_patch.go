package model

type KubernetesClusterPatch struct {
	// Shows if the Kubernetes cluster is accessible by the Bitmovin Agent
	Connected *bool `json:"connected,omitempty"`
}

