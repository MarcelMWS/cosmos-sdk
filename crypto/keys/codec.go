package keys

import (
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"

	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/codec"
	"repo.mwaysolutions.com/blockscape/gaia-yubihsm/crypto/keys/hd"
)

var cdc *codec.Codec

func init() {
	cdc = codec.New()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*Info)(nil), nil)
	cdc.RegisterConcrete(hd.BIP44Params{}, "crypto/keys/hd/BIP44Params", nil)
	cdc.RegisterConcrete(localInfo{}, "crypto/keys/localInfo", nil)
	cdc.RegisterConcrete(ledgerInfo{}, "crypto/keys/ledgerInfo", nil)
	cdc.RegisterConcrete(offlineInfo{}, "crypto/keys/offlineInfo", nil)
	cdc.RegisterConcrete(multiInfo{}, "crypto/keys/multiInfo", nil)
	cdc.Seal()
}
