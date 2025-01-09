package types

// PlayerID represents a id for an Player
type PlayerID struct {
	GenID
}

var NilPlayerID = PlayerID{GenID: NilID}
