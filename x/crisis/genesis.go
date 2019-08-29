package crisis

import (
	sdk "repo.mwaysolutions.com/blockscape/gaia-yubihsm/types"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/crisis/internal/keeper"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/crisis/internal/types"
)

// new crisis genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	keeper.SetConstantFee(ctx, data.ConstantFee)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) types.GenesisState {
	constantFee := keeper.GetConstantFee(ctx)
	return types.NewGenesisState(constantFee)
}
