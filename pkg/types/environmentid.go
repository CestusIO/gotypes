package types
//EnvironmentID represents an id for an environment
type EnvironmentID struct {
	GenID
}

var NilEnvironmentID = EnvironmentID{GenID: NilID}