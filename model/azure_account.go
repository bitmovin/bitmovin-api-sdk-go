package model


// AzureAccount model
type AzureAccount struct {
    // Name of the resource. Can be freely chosen by the user.
    Name *string `json:"name,omitempty"`
    // Description of the resource. Can be freely chosen by the user.
    Description *string `json:"description,omitempty"`
    // Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
    CreatedAt *DateTime `json:"createdAt,omitempty"`
    // Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
    ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
    // User-specific meta data. This can hold anything.
    CustomData *map[string]interface{} `json:"customData,omitempty"`
    // Id of the resource (required)
    Id *string `json:"id,omitempty"`
    // Your Azure Subscription ID (The ID of your subscription where you intend to run the Encoding VMs) (required)
    SubscriptionId *string `json:"subscriptionId,omitempty"`
    // The name of the resource group where you intend to run the Encoding VMs (required)
    ResourceGroupId *string `json:"resourceGroupId,omitempty"`
    // The ID of your Active Directory where your subscription runs in (required)
    TenantId *string `json:"tenantId,omitempty"`
}



