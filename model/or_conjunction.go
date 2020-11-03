package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
)

// OrConjunction model
type OrConjunction struct {
    // Array to perform the AND/OR evaluation on
    Conditions []AbstractCondition `json:"conditions,omitempty"`
}

func (m OrConjunction) ConditionType() ConditionType {
    return ConditionType_OR
}
// UnmarshalJSON unmarshals model OrConjunction from a JSON structure
func (m *OrConjunction) UnmarshalJSON(raw []byte) error {
    var data struct {
        
        Conditions json.RawMessage `json:"conditions"`
    }

    buf := bytes.NewBuffer(raw)
    dec := json.NewDecoder(buf)
    dec.UseNumber()

    if err := dec.Decode(&data); err != nil {
        return err
    }

    var result OrConjunction
    
    
    allOfConditions, err := UnmarshalAbstractConditionSlice(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
    if err != nil && err != io.EOF {
        return err
    }

    result.Conditions = allOfConditions

    *m = result

    return nil
}

func (m OrConjunction) MarshalJSON() ([]byte, error) {
    type M OrConjunction
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "OR"

    return json.Marshal(x)
}


