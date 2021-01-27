package client

import (
	govclient "github.com/fastock/fastock-chain/x/gov/client"
	"github.com/fastock/fastock-chain/x/params/client/cli"
	"github.com/fastock/fastock-chain/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler in cmsdk
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
