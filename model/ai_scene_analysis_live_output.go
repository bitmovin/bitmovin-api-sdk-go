package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AiSceneAnalysisLiveOutput model
type AiSceneAnalysisLiveOutput struct {
	// ID of an existing Encoding Output owned by the organization. Set either this property or `output`, but not both.
	OutputId *string `json:"outputId,omitempty"`
	// Inline definition of a concrete, publicly creatable Encoding Output to create synchronously. Only properties defined by the selected concrete Output type are accepted; internal types and properties are not supported. Deprecated properties that remain supported by the Encoding Output creation API are accepted. Set either this property or `outputId`, but not both. Put ACL entries on the destination-level `acl` property, not in this resource definition. The created Output is an ordinary reusable Encoding resource and is not automatically deleted with the Live Analysis or after provisioning failure.
	Output *Output `json:"output,omitempty"`
	// Subdirectory where files are written. This destination setting is not part of the inline Output resource definition. (required)
	OutputPath *string `json:"outputPath,omitempty"`
	// Determines accessibility of files written to this destination. Only applies to Output types that support ACLs. Defaults to PUBLIC_READ if the list is empty.
	Acl []AclEntry `json:"acl,omitempty"`
}

// UnmarshalJSON unmarshals model AiSceneAnalysisLiveOutput from a JSON structure
func (m *AiSceneAnalysisLiveOutput) UnmarshalJSON(raw []byte) error {
	var data struct {
		OutputId   *string         `json:"outputId"`
		Output     json.RawMessage `json:"output"`
		OutputPath *string         `json:"outputPath"`
		Acl        []AclEntry      `json:"acl"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AiSceneAnalysisLiveOutput

	result.OutputId = data.OutputId
	result.OutputPath = data.OutputPath
	result.Acl = data.Acl

	allOfOutput, err := UnmarshalOutput(bytes.NewBuffer(data.Output), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Output = &allOfOutput

	*m = result

	return nil
}
