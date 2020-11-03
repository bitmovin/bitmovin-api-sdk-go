package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
)

// DefaultManifestOrCondition model
type DefaultManifestOrCondition struct {
    // Array to perform the OR evaluation on. This conditions evaluates to true if at least one sub condition evaluates to true. 
    Conditions []DefaultManifestCondition `json:"conditions,omitempty"`
}

func (m DefaultManifestOrCondition) ConditionType() ConditionType {
    return ConditionType_OR
}
// UnmarshalJSON unmarshals model DefaultManifestOrCondition from a JSON structure
func (m *DefaultManifestOrCondition) UnmarshalJSON(raw []byte) error {
    var data struct {
        
        Conditions json.RawMessage `json:"conditions"`
    }

    buf := bytes.NewBuffer(raw)
    dec := json.NewDecoder(buf)
    dec.UseNumber()

    if err := dec.Decode(&data); err != nil {
        return err
    }

    var result DefaultManifestOrCondition
    
    
    allOfConditions, err := UnmarshalDefaultManifestConditionSlice(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
    if err != nil && err != io.EOF {
        return err
    }

    result.Conditions = allOfConditions

    *m = result

    return nil
}

func (m DefaultManifestOrCondition) MarshalJSON() ([]byte, error) {
    type M DefaultManifestOrCondition
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "OR"

    return json.Marshal(x)
}


