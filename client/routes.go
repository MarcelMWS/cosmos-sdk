package client

import (
	"github.com/gorilla/mux"

	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/client/context"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	RegisterRPCRoutes(cliCtx, r)
}
