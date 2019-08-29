package types

import (
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/codec"
	sdk "repo.mwaysolutions.com/blockscape/gaia-yubihsm/types"
	authtypes "repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/auth/types"
	stakingtypes "repo.mwaysolutions.com/blockscape/gaia-yubihsm/x/staking/types"
)

// ModuleCdc defines a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

// TODO: abstract genesis transactions registration back to staking
// required for genesis transactions
func init() {
	ModuleCdc = codec.New()
	stakingtypes.RegisterCodec(ModuleCdc)
	authtypes.RegisterCodec(ModuleCdc)
	sdk.RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
