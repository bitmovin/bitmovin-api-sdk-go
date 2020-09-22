package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// DefaultDashManifestPeriod model
type DefaultDashManifestPeriod struct {
	// List the encoding ids for which the conditions should apply
	EncodingIds []string `json:"encodingIds,omitempty"`
	// Adds an adaption set for every item to the period
	AdaptationSets []DefaultManifestCondition `json:"adaptationSets,omitempty"`
}

// UnmarshalJSON unmarshals model DefaultDashManifestPeriod from a JSON structure
func (m *DefaultDashManifestPeriod) UnmarshalJSON(raw []byte) error {
	var data struct {
		EncodingIds    []string        `json:"encodingIds"`
		AdaptationSets json.RawMessage `json:"adaptationSets"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result DefaultDashManifestPeriod

	result.EncodingIds = data.EncodingIds

	allOfAdaptationSets, err := UnmarshalDefaultManifestConditionSlice(bytes.NewBuffer(data.AdaptationSets), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.AdaptationSets = allOfAdaptationSets

	*m = result

	return nil
}
