package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// EncodingInfrastructureAPI intermediary API object with no endpoints
type EncodingInfrastructureAPI struct {
	apiClient *apiclient.APIClient

	// Kubernetes communicates with '/encoding/infrastructure/kubernetes' endpoints
	Kubernetes *EncodingInfrastructureKubernetesAPI
	// Aws communicates with '/encoding/infrastructure/aws' endpoints
	Aws *EncodingInfrastructureAwsAPI
	// StaticIps communicates with '/encoding/infrastructure/static-ips' endpoints
	StaticIps *EncodingInfrastructureStaticIpsAPI
	// Azure communicates with '/encoding/infrastructure/azure' endpoints
	Azure *EncodingInfrastructureAzureAPI
	// Gce communicates with '/encoding/infrastructure/gce' endpoints
	Gce *EncodingInfrastructureGceAPI
	// PrewarmedEncoderPools communicates with '/encoding/infrastructure/prewarmed-encoder-pools' endpoints
	PrewarmedEncoderPools *EncodingInfrastructurePrewarmedEncoderPoolsAPI
}

// NewEncodingInfrastructureAPI constructor for EncodingInfrastructureAPI that takes options as argument
func NewEncodingInfrastructureAPI(options ...apiclient.APIClientOption) (*EncodingInfrastructureAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInfrastructureAPIWithClient(apiClient), nil
}

// NewEncodingInfrastructureAPIWithClient constructor for EncodingInfrastructureAPI that takes an APIClient as argument
func NewEncodingInfrastructureAPIWithClient(apiClient *apiclient.APIClient) *EncodingInfrastructureAPI {
	a := &EncodingInfrastructureAPI{apiClient: apiClient}
	a.Kubernetes = NewEncodingInfrastructureKubernetesAPIWithClient(apiClient)
	a.Aws = NewEncodingInfrastructureAwsAPIWithClient(apiClient)
	a.StaticIps = NewEncodingInfrastructureStaticIpsAPIWithClient(apiClient)
	a.Azure = NewEncodingInfrastructureAzureAPIWithClient(apiClient)
	a.Gce = NewEncodingInfrastructureGceAPIWithClient(apiClient)
	a.PrewarmedEncoderPools = NewEncodingInfrastructurePrewarmedEncoderPoolsAPIWithClient(apiClient)

	return a
}
