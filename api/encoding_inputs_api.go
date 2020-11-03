package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsAPI communicates with '/encoding/inputs' endpoints
type EncodingInputsAPI struct {
    apiClient *apiclient.APIClient

    // Type communicates with '/encoding/inputs/{input_id}/type' endpoints
    Type *EncodingInputsTypeAPI
    // Rtmp communicates with '/encoding/inputs/rtmp' endpoints
    Rtmp *EncodingInputsRtmpAPI
    // RedundantRtmp communicates with '/encoding/inputs/redundant-rtmp' endpoints
    RedundantRtmp *EncodingInputsRedundantRtmpAPI
    // S3 communicates with '/encoding/inputs/s3' endpoints
    S3 *EncodingInputsS3API
    // S3RoleBased communicates with '/encoding/inputs/s3-role-based' endpoints
    S3RoleBased *EncodingInputsS3RoleBasedAPI
    // GenericS3 communicates with '/encoding/inputs/generic-s3' endpoints
    GenericS3 *EncodingInputsGenericS3API
    // Local communicates with '/encoding/inputs/local' endpoints
    Local *EncodingInputsLocalAPI
    // Gcs communicates with '/encoding/inputs/gcs' endpoints
    Gcs *EncodingInputsGcsAPI
    // GcsServiceAccount communicates with '/encoding/inputs/gcs-service-account' endpoints
    GcsServiceAccount *EncodingInputsGcsServiceAccountAPI
    // Azure communicates with '/encoding/inputs/azure' endpoints
    Azure *EncodingInputsAzureAPI
    // Ftp communicates with '/encoding/inputs/ftp' endpoints
    Ftp *EncodingInputsFtpAPI
    // Sftp communicates with '/encoding/inputs/sftp' endpoints
    Sftp *EncodingInputsSftpAPI
    // Http communicates with '/encoding/inputs/http' endpoints
    Http *EncodingInputsHttpAPI
    // Https communicates with '/encoding/inputs/https' endpoints
    Https *EncodingInputsHttpsAPI
    // Aspera communicates with '/encoding/inputs/aspera' endpoints
    Aspera *EncodingInputsAsperaAPI
    // AkamaiNetstorage communicates with '/encoding/inputs/akamai-netstorage' endpoints
    AkamaiNetstorage *EncodingInputsAkamaiNetstorageAPI
    // Srt communicates with '/encoding/inputs/srt' endpoints
    Srt *EncodingInputsSrtAPI
    // Tcp communicates with '/encoding/inputs/tcp' endpoints
    Tcp *EncodingInputsTcpAPI
    // Udp communicates with '/encoding/inputs/udp' endpoints
    Udp *EncodingInputsUdpAPI
    // UdpMulticast communicates with '/encoding/inputs/udp-multicast' endpoints
    UdpMulticast *EncodingInputsUdpMulticastAPI
    // Zixi communicates with '/encoding/inputs/zixi' endpoints
    Zixi *EncodingInputsZixiAPI

}

// NewEncodingInputsAPI constructor for EncodingInputsAPI that takes options as argument
func NewEncodingInputsAPI(options ...apiclient.APIClientOption) (*EncodingInputsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsAPIWithClient(apiClient), nil
}

// NewEncodingInputsAPIWithClient constructor for EncodingInputsAPI that takes an APIClient as argument
func NewEncodingInputsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAPI {
    a := &EncodingInputsAPI{apiClient: apiClient}
    a.Type = NewEncodingInputsTypeAPIWithClient(apiClient)
    a.Rtmp = NewEncodingInputsRtmpAPIWithClient(apiClient)
    a.RedundantRtmp = NewEncodingInputsRedundantRtmpAPIWithClient(apiClient)
    a.S3 = NewEncodingInputsS3APIWithClient(apiClient)
    a.S3RoleBased = NewEncodingInputsS3RoleBasedAPIWithClient(apiClient)
    a.GenericS3 = NewEncodingInputsGenericS3APIWithClient(apiClient)
    a.Local = NewEncodingInputsLocalAPIWithClient(apiClient)
    a.Gcs = NewEncodingInputsGcsAPIWithClient(apiClient)
    a.GcsServiceAccount = NewEncodingInputsGcsServiceAccountAPIWithClient(apiClient)
    a.Azure = NewEncodingInputsAzureAPIWithClient(apiClient)
    a.Ftp = NewEncodingInputsFtpAPIWithClient(apiClient)
    a.Sftp = NewEncodingInputsSftpAPIWithClient(apiClient)
    a.Http = NewEncodingInputsHttpAPIWithClient(apiClient)
    a.Https = NewEncodingInputsHttpsAPIWithClient(apiClient)
    a.Aspera = NewEncodingInputsAsperaAPIWithClient(apiClient)
    a.AkamaiNetstorage = NewEncodingInputsAkamaiNetstorageAPIWithClient(apiClient)
    a.Srt = NewEncodingInputsSrtAPIWithClient(apiClient)
    a.Tcp = NewEncodingInputsTcpAPIWithClient(apiClient)
    a.Udp = NewEncodingInputsUdpAPIWithClient(apiClient)
    a.UdpMulticast = NewEncodingInputsUdpMulticastAPIWithClient(apiClient)
    a.Zixi = NewEncodingInputsZixiAPIWithClient(apiClient)

    return a
}

// Get Input Details
func (api *EncodingInputsAPI) Get(inputId string) (*model.Input, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.Input
    err := api.apiClient.Get("/encoding/inputs/{input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all Inputs
func (api *EncodingInputsAPI) List(queryParams ...func(*EncodingInputsAPIListQueryParams)) (*pagination.InputsListPagination, error) {
    queryParameters := &EncodingInputsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.InputsListPagination
    err := api.apiClient.Get("/encoding/inputs", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingInputsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


