package server

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/abci/server"

	"github.com/certusone/aiakos"
	tcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	// "github.com/tendermint/tendermint/crypto/ed25519"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	// pvm "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	"os"
	"strconv"
)

// Tendermint full-node start flags
const (
	flagWithTendermint = "with-tendermint"
	flagAddress        = "address"
	flagTraceStore     = "trace-store"
	flagPruning        = "pruning"
	FlagMinGasPrices   = "minimum-gas-prices"
)

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func StartCmd(ctx *Context, appCreator AppCreator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Run the full node",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !viper.GetBool(flagWithTendermint) {
				ctx.Logger.Info("Starting ABCI without Tendermint")
				return startStandAlone(ctx, appCreator)
			}

			ctx.Logger.Info("Starting ABCI with Tendermint")

			_, err := startInProcess(ctx, appCreator)
			return err
		},
	}

	// core flags for the ABCI application
	cmd.Flags().Bool(flagWithTendermint, true, "Run abci app embedded in-process with tendermint")
	cmd.Flags().String(flagAddress, "tcp://0.0.0.0:26658", "Listen address")
	cmd.Flags().String(flagTraceStore, "", "Enable KVStore tracing to an output file")
	cmd.Flags().String(flagPruning, "syncable", "Pruning strategy: syncable, nothing, everything")
	cmd.Flags().String(
		FlagMinGasPrices, "",
		"Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum (e.g. 0.01photino;0.0001stake)",
	)

	// add support for all Tendermint-specific command line options
	tcmd.AddNodeFlags(cmd)
	return cmd
}

func startStandAlone(ctx *Context, appCreator AppCreator) error {
	addr := viper.GetString(flagAddress)
	home := viper.GetString("home")
	traceWriterFile := viper.GetString(flagTraceStore)

	db, err := openDB(home)
	if err != nil {
		return err
	}
	traceWriter, err := openTraceWriter(traceWriterFile)
	if err != nil {
		return err
	}

	app := appCreator(ctx.Logger, db, traceWriter)

	svr, err := server.NewServer(addr, "socket", app)
	if err != nil {
		return fmt.Errorf("error creating listener: %v", err)
	}

	svr.SetLogger(ctx.Logger.With("module", "abci-server"))

	err = svr.Start()
	if err != nil {
		cmn.Exit(err.Error())
	}

	// wait forever
	cmn.TrapSignal(ctx.Logger, func() {
		// cleanup
		err = svr.Stop()
		if err != nil {
			cmn.Exit(err.Error())
		}
	})
	return nil
}

func startInProcess(ctx *Context, appCreator AppCreator) (*node.Node, error) {
	cfg := ctx.Config
	home := cfg.RootDir
	traceWriterFile := viper.GetString(flagTraceStore)

	db, err := openDB(home)
	if err != nil {
		return nil, err
	}
	traceWriter, err := openTraceWriter(traceWriterFile)
	if err != nil {
		return nil, err
	}

	app := appCreator(ctx.Logger, db, traceWriter)

	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		return nil, err
	}

	UpgradeOldPrivValFile(cfg)
	if os.Getenv("AIAKOS_URL") == "" {
		return nil, errors.New("no Aiakos hsm url specified. Please set AIAKOS_URL in the format host:port")
	}
	aiakosUrl := os.Getenv("AIAKOS_URL")
	if os.Getenv("AIAKOS_SIGNING_KEY") == "" {
		return nil, errors.New("no Aiakos signing key ID specified. Please set AIAKOS_SIGNING_KEY")
	}
	aiakosSigningKey, err := strconv.ParseUint(os.Getenv("AIAKOS_SIGNING_KEY"), 10, 16)
	if err != nil {
		return nil, errors.New("invalid Aiakos signing key ID.")
	}
	if os.Getenv("AIAKOS_AUTH_KEY") == "" {
		return nil, errors.New("no Aiakos auth key ID specified. Please set AIAKOS_AUTH_KEY")
	}
	aiakosAuthKey, err := strconv.ParseUint(os.Getenv("AIAKOS_AUTH_KEY"), 10, 16)
	if err != nil {
		return nil, errors.New("invalid Aiakos auth key ID.")
	}
	if os.Getenv("AIAKOS_AUTH_KEY_PASSWORD") == "" {
		return nil, errors.New("no Aiakos auth key password specified. Please set AIAKOS_AUTH_KEY_PASSWORD")
	}
	aiakosAuthPassword := os.Getenv("AIAKOS_AUTH_KEY_PASSWORD")
	// Init Aiakos module
	hsm, err := aiakos.NewAiakosPV(aiakosUrl, uint16(aiakosSigningKey), uint16(aiakosAuthKey), aiakosAuthPassword, ctx.Logger.With("module", "aiakos"))
	if err != nil {
		return nil, err
	}
	// Start Aiakos
	err = hsm.Start()
	if err != nil {
		return nil, err
	}
	if os.Getenv("AIAKOS_IMPORT_KEY") == "TRUE" {
		ctx.Logger.Info("importing private key to Aiakos because AIAKOS_IMPORT_KEY is set.")
		// filepv := pvm.LoadOrGenFilePV
		// key := pvm.FilePVKey{PrivKey: ed25519.PrivKeyEd25519{}}
		x := []byte("A");
		err = hsm.ImportKey(uint16(aiakosSigningKey), x)
		if err != nil {
			ctx.Logger.Error("Could not import key to HSM; skipping this step since it probably already exists", "error", err)
		}
	}
	// create & start tendermint node
	tmNode, err := node.NewNode(
		cfg,
		hsm,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(cfg),
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(cfg.Instrumentation),
		ctx.Logger.With("module", "node"),
	)
	if err != nil {
		return nil, err
	}

	err = tmNode.Start()
	if err != nil {
		return nil, err
	}

	TrapSignal(func() {
		if tmNode.IsRunning() {
			_ = tmNode.Stop()
		}
	})

	// run forever (the node will not be returned)
	select {}
}