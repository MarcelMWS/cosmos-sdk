package client

import (
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/distribution/client/cli"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/distribution/client/rest"
	govclient "repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
