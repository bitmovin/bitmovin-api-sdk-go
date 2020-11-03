package model

import (
    "encoding/json"
)

// Condition model
type Condition struct {
    // The attribute that should be used for the evaluation. Valid values include, depending on the context: - HEIGHT - WIDTH - BITRATE - FPS - ASPECTRATIO - INPUTSTREAM - LANGUAGE - CHANNELFORMAT - CHANNELLAYOUT - STREAMCOUNT - AUDIOSTREAMCOUNT - VIDEOSTREAMCOUNT - DURATION - ROTATION (required)
    Attribute *string `json:"attribute,omitempty"`
    // The operator that should be used for the evaluation (required)
    Operator ConditionOperator `json:"operator,omitempty"`
    // The value that should be used for comparison (required)
    Value *string `json:"value,omitempty"`
}

func (m Condition) ConditionType() ConditionType {
    return ConditionType_CONDITION
}
func (m Condition) MarshalJSON() ([]byte, error) {
    type M Condition
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "CONDITION"

    return json.Marshal(x)
}


