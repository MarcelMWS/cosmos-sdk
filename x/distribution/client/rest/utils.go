package rest

import (
	sdk "repo.mwaysolutions.com/blockscape/gaia-yubihsm/types"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/types/rest"
)

type (
	// CommunityPoolSpendProposalReq defines a community pool spend proposal request body.
	CommunityPoolSpendProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title       string         `json:"title" yaml:"title"`
		Description string         `json:"description" yaml:"description"`
		Recipient   sdk.AccAddress `json:"recipient" yaml:"recipient"`
		Amount      sdk.Coins      `json:"amount" yaml:"amount"`
		Proposer    sdk.AccAddress `json:"proposer" yaml:"proposer"`
		Deposit     sdk.Coins      `json:"deposit" yaml:"deposit"`
	}
)
