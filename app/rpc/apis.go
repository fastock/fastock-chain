package rpc

import (
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/cosmos/cosmos-sdk/client/context"

	"../../app/crypto/ethsecp256k1"
	"../../app/rpc/backend"
	"../../app/rpc/namespaces/eth"
	"../../app/rpc/namespaces/eth/filters"
	"../../app/rpc/namespaces/net"
	"../../app/rpc/namespaces/personal"
	"../../app/rpc/namespaces/web3"
	rpctypes "../../app/rpc/types"
)

// RPC namespaces and API version
const (
	Web3Namespace     = "web3"
	EthNamespace      = "eth"
	PersonalNamespace = "personal"
	NetNamespace      = "net"

	apiVersion = "1.0"
)

// GetAPIs returns the list of all APIs from the Ethereum namespaces
func GetAPIs(clientCtx context.CLIContext, keys ...ethsecp256k1.PrivKey) []rpc.API {
	nonceLock := new(rpctypes.AddrLocker)
	backend := backend.New(clientCtx)
	ethAPI := eth.NewAPI(clientCtx, backend, nonceLock, keys...)

	return []rpc.API{
		{
			Namespace: Web3Namespace,
			Version:   apiVersion,
			Service:   web3.NewAPI(),
			Public:    true,
		},
		{
			Namespace: EthNamespace,
			Version:   apiVersion,
			Service:   ethAPI,
			Public:    true,
		},
		{
			Namespace: EthNamespace,
			Version:   apiVersion,
			Service:   filters.NewAPI(clientCtx, backend),
			Public:    true,
		},
		{
			Namespace: PersonalNamespace,
			Version:   apiVersion,
			Service:   personal.NewAPI(ethAPI),
			Public:    false,
		},
		{
			Namespace: NetNamespace,
			Version:   apiVersion,
			Service:   net.NewAPI(clientCtx),
			Public:    true,
		},
	}
}
