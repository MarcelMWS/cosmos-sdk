package cli

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types"
)

// IBC transfer flags
var (
	FlagNode1    = "node1"
	FlagNode2    = "node2"
	FlagFrom1    = "from1"
	FlagFrom2    = "from2"
	FlagChainID2 = "chain-id2"
	FlagSequence = "packet-sequence"
	FlagTimeout  = "timeout"
)

// GetTransferTxCmd returns the command to create a NewMsgTransfer transaction
func GetTransferTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [src-port] [src-channel] [dest-height] [receiver] [amount]",
		Short: "Transfer fungible token through IBC",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock)

			sender := cliCtx.GetFromAddress()
			srcPort := args[0]
			srcChannel := args[1]
			destHeight, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			// parse coin trying to be sent
			coins, err := sdk.ParseCoins(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransfer(srcPort, srcChannel, uint64(destHeight), coins, sender, args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}

// GetTroughputTxCmd returns the command to create a NewMsgTransfer transaction
func GetTroughputTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "throughput [src-port] [src-channel] [dest-height] [receiver] [amount]",
		Short: "start throughput txs for GoZ",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigName("throughput")
			viper.SetConfigType("yaml")

			//Set defaults
			viper.SetDefault("paths", map[string]string{"chain-id": "", "client-id": "", "connection-id": "", "channel-id": "abcdefghijk", "port-id": "transfer", "coin": "gain"})
			viper.SetDefault("pathd", map[string]string{"chain-id": "", "client-id": "", "connection-id": "", "channel-id": "abcdefghijk", "port-id": "transfer", "coin": "doubloons"})
			viper.SetDefault("rec_address", "cosmos1kq7yuwmanzj60p6ng0p9aj0ttsyd4ynaqvj36n")
			viper.SetDefault("dest-height", "0")
			viper.SetDefault("amount", "1")
			viper.SetDefault("node", "tcp://ibc.blockscape.:26657")
			viper.SetDefault("chain-id", "abchain")
			viper.SetDefault("send_accounts", map[int]string{0: "cosmos12f3pu4cn9frg5t3pn2acwywn8z8ds9uqfhxeml", 1: "", 2: "", 3: "", 4: "", 5: "", 6: ""})

			//safe config
			viper.SafeWriteConfig()

			//read config
			err := viper.ReadInConfig() // Find and read the config file
			if err != nil {             // Handle errors reading the config file
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock).WithGenerateOnly(true).WithChainID(viper.GetString("chain-id"))

			// get values
			// log.Println(viper.Get("pathd.client-id"))

			sender, err := sdk.AccAddressFromBech32(viper.GetString("send_accounts.0"))
			srcPort := viper.GetString("paths.port-id")
			srcChannel := viper.GetString("paths.channel-id")
			destHeight, err := strconv.Atoi(viper.GetString("dest-height"))
			if err != nil {
				return err
			}

			// parse coin trying to be sent
			coins, err := sdk.ParseCoins(viper.GetString("amount") + viper.GetString("paths.port-id") + "/" + viper.GetString("pathd.channel-id") + "/" + viper.GetString("paths.coin"))
			if err != nil {
				return err
			}
			log.Println("Coins: ", coins)

			//edit sender
			msg := types.NewMsgTransfer(srcPort, srcChannel, uint64(destHeight), coins, sender, viper.GetString("rec_address"))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}
