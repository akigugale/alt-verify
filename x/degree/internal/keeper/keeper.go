package keeper

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/akigugale/alt-verify/x/degree/internal/types"
)

// Keeper of the degree store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

func (k Keeper) SetDegree(ctx sdk.Context, degree types.Degree) {
	if degree.aadhar
}