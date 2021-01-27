package evm

import (
	"github.com/fastock/fastock-chain/x/evm/keeper"
	"github.com/fastock/fastock-chain/x/evm/types"
)

// nolint
const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	DefaultParamspace = types.DefaultParamspace
)

// nolint
var (
	NewKeeper = keeper.NewKeeper
	TxDecoder = types.TxDecoder
)

//nolint
type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
