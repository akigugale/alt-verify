package keeper

import (
	"fmt"
	"github.com/akigugale/alt-verify/x/degree/internal/types"
)


func SetDegree(degree types.Degree) {
	fmt.Printf("%+v", degree)
}


types.Degree d = types.NewDegree(university="VIT", universityName="Vellore institute of technology", aadhar=237410872819)

SetDegree(d)
