package types

// EOSProductID represents an id for a Product
type EOSProductID struct {
	GenID
}

var NilEOSProductID = EOSProductID{GenID: NilID}
