package types

// EOSDeploymentID represents an id for a Deployment
type EOSDeploymentID struct {
	GenID
}

var NilEOSDeploymentID = EOSDeploymentID{GenID: NilID}