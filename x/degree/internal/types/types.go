package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Degree struct {
	University sdk.AccAddress `json:"university"`
	UniversityName string `json:"university_name"`
	Degree string `json:"degree"`
	Subject string `json:"subject"`
	Duration int `json:"duration"` // in months
	Batch string `json:"batch"`
	StudentName string `json:"student_name"`
	AadharNo int64 `json:"aadhar"`
	RollNo string `json:"roll_no"`
	Gpa float32 `json:"gpa"`
	Status bool `json:"status"`
	Meta string `json:"meta,omitempty"`
	Created time.Time `json:"created"`
}

func NewDegree (university sdk.AccAddress, universityName string, degree string, subject string, duration int, batch string, studentName string, rollNo string, gpa float32, status bool, aadharNo int64, meta string, created time.Time) Degree {
	return Degree {
		University: university,
		UniversityName: universityName,
		Degree: degree,
		Subject: subject,
		Duration: duration,
		Batch: batch,
		StudentName: studentName,
		AadharNo: aadharNo,
		RollNo: rollNo,
		Gpa: gpa,
		Status: status,
		Meta: meta,
		Created: created
	}
}


// fmt.Stringer
func (d Degree) String() string {
	return strings.TrimSpace("%+v", d)
}

// TODO: add if degree of same aadhar no found within past x years of this degree