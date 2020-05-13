package keeper

import (
	"fmt"
	"bytes"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/akigugale/alt-verify/x/degree/types"

)

// Keeper of the degree store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a alt-verify keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc     : cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetDegree returns degrees for a particular student (AccAddresss)
func (k Keeper) GetDegree(ctx sdk.Context, student sdk.AccAddress) (types.Degree, error) {
	store := ctx.KVStore(k.storeKey)
	var degree types.Degree
	byteKey := []byte(student)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &degree)
	if err != nil {
		return degree, err
	}
	return degree, nil
}

// Creates a student --> degree key value pair NOTE: what if KV pair already exists
func (k Keeper) SetDegree(ctx sdk.Context, degree types.Degree) {
	student := degree.Student
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(degree)
	key := []byte(student)
	store.Set(key, bz)
}

// Deletes a degree corresponding to student
func (k Keeper) DeleteDegree(ctx sdk.Context, student sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(student))
}

// Get an iterator over all students in which the keys are the names and the values are the degree
func (k Keeper) GetDegreeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the student is present in the store or not
func (k Keeper) IsStudentPresent(ctx sdk.Context, student sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(student))
}


func (k Keeper) GetDegreeOfUni(ctx sdk.Context, creator sdk.AccAddress) ([]string, error) {
	// store := ctx.KVStore(k.storeKey)

	var degreeList []string
    iterator := k.GetDegreeIterator(ctx)
	// var degree types.Degree
    for ; iterator.Valid(); iterator.Next() {
		degree, err := k.GetDegree(ctx, iterator.Key())
		if err != nil {
			continue;
		}
		if bytes.Compare(degree.Creator, creator) == 0 {
			degreeList = append(degreeList, string(degree.Student))
		}
	}

    if len(degreeList) == 0 {
        return degreeList, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown Uni query endpoint")
    }

    return degreeList, nil
}