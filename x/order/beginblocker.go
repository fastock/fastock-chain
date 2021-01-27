package order

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/fastock/fastock-chain/x/common/perf"
	"github.com/fastock/fastock-chain/x/order/keeper"
	"github.com/fastock/fastock-chain/x/order/types"
	//"github.com/fastock/fastock-chain/x/common/version"
)

// BeginBlocker runs the logic of BeginBlocker with version 0.
// BeginBlocker resets keeper cache.
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
