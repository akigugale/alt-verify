package types

import "strings"

// Query endpoints supported by the degree querier
const (
	QueryListDegrees = "list"
	QueryGetDegree = "get"
	QueryGetDegreesOfUni = "listuni"
)



// QueryResList Queries Result Payload for a query
type QueryResDegrees []string

// implement fmt.Stringer
func (n QueryResDegrees) String() string {
	return strings.Join(n[:], "\n")
}

