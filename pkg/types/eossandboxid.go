package types
//EOSSandboxID represents an id for a Sandbox
type EOSSandboxID struct {
	GenID
}

var NilEOSSandboxID = EOSSandboxID{GenID: NilID}