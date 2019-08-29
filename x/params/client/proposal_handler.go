package client

import (
	govclient "repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/gov/client"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/params/client/cli"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
