package continuousauction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"../../x/order/keeper"
)

// nolint
type CaEngine struct {
}

// nolint
func (e *CaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	// TODO
}
