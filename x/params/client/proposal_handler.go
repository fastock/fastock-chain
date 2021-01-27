package client

import (
	govclient "../../x/gov/client"
	"../../x/params/client/cli"
	"../../x/params/client/rest"
)

// ProposalHandler is the param change proposal handler in cmsdk
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
