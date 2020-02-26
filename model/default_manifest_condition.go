package model

type DefaultManifestCondition interface {
    ConditionType() ConditionType
}

type BaseDefaultManifestCondition struct {
    Type ConditionType `json:"type"`
}

func (b BaseDefaultManifestCondition) ConditionType() ConditionType {
    return b.Type
}

