package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "blockchain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "blockchain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "blockchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "blockchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "blockchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "blockchain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "blockchain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "blockchain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "blockchain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
