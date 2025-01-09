package types

// ApplicationID represents a id for an application
type ApplicationID struct {
	GenID
}

var NilApplicationID = ApplicationID{GenID: NilID}
