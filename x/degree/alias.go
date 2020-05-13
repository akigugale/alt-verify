package degree

import (
	"github.com/akigugale/alt-verify/x/degree/keeper"
	"github.com/akigugale/alt-verify/x/degree/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc     = types.ModuleCdc

	// Msg aliases
	NewMsgCreateDegree= types.NewMsgCreateDegree
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	// Msg
	MsgCreateDegree = types.MsgCreateDegree
)
