package client

import (
	"../../x/distribution/client/cli"
	"../../x/distribution/client/rest"
	govclient "../../x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
