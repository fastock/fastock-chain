package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/fastock/fastock-chain/x/ammswap"
	ammswaptypes "github.com/fastock/fastock-chain/x/ammswap/types"
	"github.com/fastock/fastock-chain/x/dex"
	farmtypes "github.com/fastock/fastock-chain/x/farm/types"
	"github.com/fastock/fastock-chain/x/order"
	"github.com/fastock/fastock-chain/x/token"
	"github.com/willf/bitset"
)

type OrderKeeper interface {
	GetOrder(ctx sdk.Context, orderID string) *order.Order
	GetUpdatedOrderIDs() []string
	GetBlockOrderNum(ctx sdk.Context, blockHeight int64) int64
	GetBlockMatchResult() *order.BlockMatchResult
	GetLastPrice(ctx sdk.Context, product string) sdk.Dec
	GetBestBidAndAsk(ctx sdk.Context, product string) (sdk.Dec, sdk.Dec)
	GetUpdatedDepthbookKeys() []string
	GetDepthBookCopy(product string) *order.DepthBook
	GetProductPriceOrderIDs(key string) []string
	GetTxHandlerMsgResult() []bitset.BitSet
}

type TokenKeeper interface {
	GetFeeDetailList() []*token.FeeDetail
	GetCoinsInfo(ctx sdk.Context, addr sdk.AccAddress) token.CoinsInfo
}

type AccountKeeper interface {
	SetObserverKeeper(observer auth.ObserverI)
}

type DexKeeper interface {
	GetTokenPairs(ctx sdk.Context) []*dex.TokenPair
	SetObserverKeeper(keeper dex.StreamKeeper)
}

// SwapKeeper expected swap keeper
type SwapKeeper interface {
	SetObserverKeeper(k ammswaptypes.BackendKeeper)
	GetSwapTokenPairs(ctx sdk.Context) []ammswap.SwapTokenPair
}

// FarmKeeper expected farm keeper
type FarmKeeper interface {
	SetObserverKeeper(k farmtypes.BackendKeeper)
}
