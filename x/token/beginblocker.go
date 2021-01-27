package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fastock/fastock-chain/x/common/perf"
	"github.com/fastock/fastock-chain/x/token/types"
)

// BeginBlocker is called when dapp handles with abci::BeginBlock
func beginBlocker(ctx sdk.Context, keeper Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
