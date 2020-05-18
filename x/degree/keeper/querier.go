package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/akigugale/alt-verify/x/degree/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for degree clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListDegrees:
			return listDegrees(ctx, k)
		case types.QueryGetDegree:
			return getDegree(ctx, path[1:], k)
		case types.QueryGetDegreesOfUni:
			return getDegreesOfUni(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown degree query endpoint")
		}
	}
}


// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()

func listDegrees(ctx sdk.Context, k Keeper) ([]byte, error) {
	var degreeList types.QueryResDegrees

	iterator := k.GetDegreeIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		degree := iterator.Key()
		degreeList = append(degreeList, string(degree))
	}

	res, err := codec.MarshalJSONIndent(k.cdc, degreeList)

	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getDegree(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	student := path[0]
	degree, err := k.GetDegree(ctx, student)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, degree)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getDegreesOfUni(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	uni := path[0]
	var degreeList types.QueryResDegrees
	degreeList, err := k.GetDegreeOfUni(ctx, sdk.AccAddress(uni))
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, degreeList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}