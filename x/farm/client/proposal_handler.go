package client

import (
	"github.com/fastock/fastock-chain/x/farm/client/cli"
	"github.com/fastock/fastock-chain/x/farm/client/rest"
	govcli "github.com/fastock/fastock-chain/x/gov/client"
)

var (
	// ManageWhiteListProposalHandler alias gov NewProposalHandler
	ManageWhiteListProposalHandler = govcli.NewProposalHandler(cli.GetCmdManageWhiteListProposal, rest.ManageWhiteListProposalRESTHandler)
)
