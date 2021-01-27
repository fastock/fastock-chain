// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/fastock/fastock-chain/x/distribution/keeper
// ALIASGEN: github.com/fastock/fastock-chain/x/distribution/types
// ALIASGEN: github.com/fastock/fastock-chain/x/distribution/client
package distribution

import (
	"github.com/fastock/fastock-chain/x/distribution/client"
	"github.com/fastock/fastock-chain/x/distribution/keeper"
	"github.com/fastock/fastock-chain/x/distribution/types"
)

const (
	ModuleName                  = types.ModuleName
	StoreKey                    = types.StoreKey
	RouterKey                   = types.RouterKey
	QuerierRoute                = types.QuerierRoute
	QueryParams                 = types.QueryParams
	QueryValidatorCommission    = types.QueryValidatorCommission
	QueryWithdrawAddr           = types.QueryWithdrawAddr
	ParamWithdrawAddrEnabled    = types.ParamWithdrawAddrEnabled
	DefaultParamspace           = types.DefaultParamspace
)

var (
	// functions aliases
	RegisterInvariants                       = keeper.RegisterInvariants
	NewKeeper                                = keeper.NewKeeper
	GetDelegatorWithdrawInfoAddress          = types.GetDelegatorWithdrawInfoAddress
	GetValidatorAccumulatedCommissionAddress = types.GetValidatorAccumulatedCommissionAddress
	GetDelegatorWithdrawAddrKey              = types.GetDelegatorWithdrawAddrKey
	GetValidatorAccumulatedCommissionKey     = types.GetValidatorAccumulatedCommissionKey
	NewQuerier                               = keeper.NewQuerier
	RegisterCodec                            = types.RegisterCodec
	ErrNilDelegatorAddr                      = types.ErrNilDelegatorAddr
	ErrNoValidatorCommission                 = types.ErrNoValidatorCommission
	ErrSetWithdrawAddrDisabled               = types.ErrSetWithdrawAddrDisabled
	InitialFeePool                           = types.InitialFeePool
	NewGenesisState                          = types.NewGenesisState
	DefaultGenesisState                      = types.DefaultGenesisState
	ValidateGenesis                          = types.ValidateGenesis
	NewMsgSetWithdrawAddress                 = types.NewMsgSetWithdrawAddress
	NewMsgWithdrawValidatorCommission        = types.NewMsgWithdrawValidatorCommission
	NewQueryValidatorCommissionParams        = types.NewQueryValidatorCommissionParams
	NewQueryDelegatorWithdrawAddrParams      = types.NewQueryDelegatorWithdrawAddrParams
	InitialValidatorAccumulatedCommission    = types.InitialValidatorAccumulatedCommission

	// variable aliases
	FeePoolKey                           = types.FeePoolKey
	ProposerKey                          = types.ProposerKey
	DelegatorWithdrawAddrPrefix          = types.DelegatorWithdrawAddrPrefix
	ValidatorAccumulatedCommissionPrefix = types.ValidatorAccumulatedCommissionPrefix
	ModuleCdc                            = types.ModuleCdc
	EventTypeSetWithdrawAddress          = types.EventTypeSetWithdrawAddress
	EventTypeCommission                  = types.EventTypeCommission
	EventTypeWithdrawCommission          = types.EventTypeWithdrawCommission
	EventTypeProposerReward              = types.EventTypeProposerReward
	AttributeKeyWithdrawAddress          = types.AttributeKeyWithdrawAddress
	AttributeKeyValidator                = types.AttributeKeyValidator
	AttributeValueCategory               = types.AttributeValueCategory
	ProposalHandler                      = client.ProposalHandler
)

type (
	Hooks                                = keeper.Hooks
	Keeper                               = keeper.Keeper
	DelegatorWithdrawInfo                = types.DelegatorWithdrawInfo
	ValidatorAccumulatedCommissionRecord = types.ValidatorAccumulatedCommissionRecord
	GenesisState                         = types.GenesisState
	MsgSetWithdrawAddress                = types.MsgSetWithdrawAddress
	MsgWithdrawValidatorCommission       = types.MsgWithdrawValidatorCommission
	QueryValidatorCommissionParams       = types.QueryValidatorCommissionParams
	QueryDelegatorWithdrawAddrParams     = types.QueryDelegatorWithdrawAddrParams
	ValidatorAccumulatedCommission       = types.ValidatorAccumulatedCommission
)
