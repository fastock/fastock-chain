package client

import (
	"github.com/fastock/fastock-chain/x/distribution/client/cli"
	"github.com/fastock/fastock-chain/x/distribution/client/rest"
	govclient "github.com/fastock/fastock-chain/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
