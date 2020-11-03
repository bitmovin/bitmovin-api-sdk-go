package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
)

// AndConjunction model
type AndConjunction struct {
    // Array to perform the AND/OR evaluation on
    Conditions []AbstractCondition `json:"conditions,omitempty"`
}

func (m AndConjunction) ConditionType() ConditionType {
    return ConditionType_AND
}
// UnmarshalJSON unmarshals model AndConjunction from a JSON structure
func (m *AndConjunction) UnmarshalJSON(raw []byte) error {
    var data struct {
        
        Conditions json.RawMessage `json:"conditions"`
    }

    buf := bytes.NewBuffer(raw)
    dec := json.NewDecoder(buf)
    dec.UseNumber()

    if err := dec.Decode(&data); err != nil {
        return err
    }

    var result AndConjunction
    
    
    allOfConditions, err := UnmarshalAbstractConditionSlice(bytes.NewBuffer(data.Conditions), bitutils.JSONConsumer())
    if err != nil && err != io.EOF {
        return err
    }

    result.Conditions = allOfConditions

    *m = result

    return nil
}

func (m AndConjunction) MarshalJSON() ([]byte, error) {
    type M AndConjunction
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "AND"

    return json.Marshal(x)
}


