package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
)

// EmailNotificationWithStreamConditionsRequest model
type EmailNotificationWithStreamConditionsRequest struct {
    // Notify when condition resolves after it was met
    Resolve *bool `json:"resolve,omitempty"`
    Emails []string `json:"emails,omitempty"`
    Muted *bool `json:"muted,omitempty"`
    Conditions *AbstractCondition `json:"conditions,omitempty"`
}

// UnmarshalJSON unmarshals model EmailNotificationWithStreamConditionsRequest from a JSON structure
func (m *EmailNotificationWithStreamConditionsRequest) UnmarshalJSON(raw []byte) error {
    var data struct {
        
        Resolve *bool `json:"resolve"`
        Emails []string `json:"emails"`
        Muted *bool `json:"muted"`
        Conditions json.RawMessage `json:"conditions"`
    }

    buf := bytes.NewBuffer(raw)
    dec := json.NewDecoder(buf)
    dec.UseNumber()

    if err := dec.Decode(&data); err != nil {
        return err
    }

    var result EmailNotificationWithStreamConditionsRequest
    
    result.Resolve = data.Resolve
    result.Emails = data.Emails
    result.Muted = data.Muted
    
    allOfConditions, err := UnmarshalAbstractCondition(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
    if err != nil && err != io.EOF {
        return err
    }

    result.Conditions = &allOfConditions

    *m = result

    return nil
}



