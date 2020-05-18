package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Degree struct {
	Creator sdk.AccAddress `json:"address" yaml:"address"` // address of the degree creator
	Student string `json:"student" yaml:"student"` // address of student
	Subject string `json:"subject" yaml:"subject"` // address of the degree creator
	Batch uint16 `json:"batch" yaml:"batch"` // year
}


// implement fmt.Stringer
func (d Degree) String() string {
	return strings.TrimSpace((fmt.Sprintf(`Creator: %s
	Student: %s
	Subject: %s
	Batch: %s`,
	d.Creator,
	d.Student,
	d.Subject,
	d.Batch)))
}

// // fmt.Stringer
// func (d Degree) String() string {
// 	return fmt.Sprintf("<%+v>", d)
// }