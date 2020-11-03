package model


// PlayReadyAdditionalInformation model
type PlayReadyAdditionalInformation struct {
    // Custom attributes that you want to add to the WRM header. This string must be valid xml. Some DRM providers require some information in the custom attributes of the msr:pro tag of the DASH manifest, otherwise the content does not play on certain devices.
    WrmHeaderCustomAttributes *string `json:"wrmHeaderCustomAttributes,omitempty"`
}



