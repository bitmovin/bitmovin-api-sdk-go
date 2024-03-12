package model

import (
	"encoding/json"
)

// StreamsLiveResponse model
type StreamsLiveResponse struct {
	// The identifier of the stream
	Id *string `json:"id,omitempty"`
	// The title of the stream
	Title *string `json:"title,omitempty"`
	// The description of the stream
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Type of the Stream contained in the StreamsResponse
	Type StreamsType `json:"type,omitempty"`
	// The streamKey of the stream
	StreamKey *string `json:"streamKey,omitempty"`
	// The life cycle of the stream
	LifeCycle   StreamsLiveLifeCycle        `json:"lifeCycle,omitempty"`
	StyleConfig *StreamsStyleConfigResponse `json:"styleConfig,omitempty"`
	// Poster URL
	PosterUrl         *string                           `json:"posterUrl,omitempty"`
	AdConfig          *StreamsAdConfigResponse          `json:"adConfig,omitempty"`
	DomainRestriction *StreamsDomainRestrictionResponse `json:"domainRestriction,omitempty"`
	// Stream trimming information
	Trimming StreamsTrimmingStatus `json:"trimming,omitempty"`
}

func (m StreamsLiveResponse) StreamsType() StreamsType {
	return StreamsType_LIVE
}
func (m StreamsLiveResponse) MarshalJSON() ([]byte, error) {
	type M StreamsLiveResponse
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "LIVE"

	return json.Marshal(x)
}
