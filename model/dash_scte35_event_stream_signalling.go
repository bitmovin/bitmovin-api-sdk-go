package model

// DashScte35EventStreamSignalling : Defines how SCTE-35 events are delivered using DASH EventStream. <ul>   <li>XML: Events will contain an XML representation of a SCTE-35 cue message as a SpliceInfoSection element as defined in SCTE-35. The schemeIdUri will be \"urn:scte:scte35:2013:xml.</li>   <li>XML_BIN: Events will contain an XML representation of a SCTE-35 cue message as a Signal.Binary element as defined in SCTE-35. The schemeIdUri will be \"urn:scte:scte35:2014:xml+bin</li> </ul>
type DashScte35EventStreamSignalling string

// List of possible DashScte35EventStreamSignalling values
const (
	DashScte35EventStreamSignalling_XML     DashScte35EventStreamSignalling = "XML"
	DashScte35EventStreamSignalling_XML_BIN DashScte35EventStreamSignalling = "XML_BIN"
)
