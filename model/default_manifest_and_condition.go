package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
)

// DefaultManifestAndCondition model
type DefaultManifestAndCondition struct {
    // Array to perform the AND evaluation on. This conditions evaluates to true if all sub conditions evaluate to true. 
    Conditions []DefaultManifestCondition `json:"conditions,omitempty"`
}

func (m DefaultManifestAndCondition) ConditionType() ConditionType {
    return ConditionType_AND
}
// UnmarshalJSON unmarshals model DefaultManifestAndCondition from a JSON structure
func (m *DefaultManifestAndCondition) UnmarshalJSON(raw []byte) error {
    var data struct {
        
        Conditions json.RawMessage `json:"conditions"`
    }

    buf := bytes.NewBuffer(raw)
    dec := json.NewDecoder(buf)
    dec.UseNumber()

    if err := dec.Decode(&data); err != nil {
        return err
    }

    var result DefaultManifestAndCondition
    
    
    allOfConditions, err := UnmarshalDefaultManifestConditionSlice(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
    if err != nil && err != io.EOF {
        return err
    }

    result.Conditions = allOfConditions

    *m = result

    return nil
}

func (m DefaultManifestAndCondition) MarshalJSON() ([]byte, error) {
    type M DefaultManifestAndCondition
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "AND"

    return json.Marshal(x)
}


