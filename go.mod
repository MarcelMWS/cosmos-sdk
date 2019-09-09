module repo.mwaysolutions.com/blockscape/cosmos-sdk-yubihsm

go 1.12

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3
	github.com/certusone/aiakos v0.3.0
	github.com/cosmos/cosmos-sdk v0.37.0
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/cosmos/ledger-cosmos-go v0.10.3
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1
	github.com/gorilla/mux v1.7.3
	github.com/mattn/go-isatty v0.0.9
	github.com/pelletier/go-toml v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/rakyll/statik v0.1.6
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20190823183015-45b1026d81ae
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.12.4
	github.com/tendermint/tendermint v0.32.3
	github.com/tendermint/tm-db v0.1.1
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/certusone/aiakos => repo.mwaysolutions.com/blockscape/aiakos-raft v0.1.0

replace github.com/cosmos/cosmos-sdk => repo.mwaysolutions.com/blockscape/cosmos-sdk-yubihsm v0.37.0-raft-test
