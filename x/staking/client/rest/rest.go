package rest

import (
	"github.com/gorilla/mux"

	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/client/context"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
