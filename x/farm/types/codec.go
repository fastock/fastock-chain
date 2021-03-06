package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "blockchain/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(MsgDestroyPool{}, "blockchain/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(MsgLock{}, "blockchain/farm/MsgLock", nil)
	cdc.RegisterConcrete(MsgUnlock{}, "blockchain/farm/MsgUnlock", nil)
	cdc.RegisterConcrete(MsgClaim{}, "blockchain/farm/MsgClaim", nil)
	cdc.RegisterConcrete(MsgProvide{}, "blockchain/farm/MsgProvide", nil)
	cdc.RegisterConcrete(ManageWhiteListProposal{}, "blockchain/farm/ManageWhiteListProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
