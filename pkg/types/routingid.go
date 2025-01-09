package types

var (
	// RoutingIDGlobal is a id for api routing to specify global services
	RoutingIDGlobal RoutingID
	// RoutingIDTitle is an id for api routing to specify a title service(s)
	RoutingIDTitle  RoutingID
)

func init() {
	Must(&RoutingIDGlobal, "cb98a312-26b5-4d90-96f3-af0f40f63291")
	Must(&RoutingIDTitle, "951638b3-e767-4d44-b5d9-7e23d190f3dd")
}

// RoutingID is an ID used in api routing
type RoutingID struct {
	GenID
}

var NilRoutingID = RoutingID{GenID: NilID}