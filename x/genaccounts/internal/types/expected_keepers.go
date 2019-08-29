package types

import (
	sdk "repo.mwaysolutions.com/blockscape/gaia-yubihsm/types"
	authexported "repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/auth/exported"
)

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	NewAccount(sdk.Context, authexported.Account) authexported.Account
	SetAccount(sdk.Context, authexported.Account)
	IterateAccounts(ctx sdk.Context, process func(authexported.Account) (stop bool))
}
