package client

import (
	"github.com/fastock/fastock-chain/x/dex/client/cli"
	"github.com/fastock/fastock-chain/x/dex/client/rest"
	govclient "github.com/fastock/fastock-chain/x/gov/client"
)

// param change proposal handler
var (
	// DelistProposalHandler alias gov NewProposalHandler
	DelistProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDelistProposal, rest.DelistProposalRESTHandler)
)
