package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"github.com/spf13/cobra"

	"github.com/akigugale/alt-verify/x/degree/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
)


// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	degreeTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	degreeTxCmd.AddCommand(flags.PostCommands(
	// Cmd's for messages
		GetCmdCreateDegree(cdc),
	)...)

	return degreeTxCmd
}

// Example:
//
// GetCmdCreateDegree is the CLI command for doing CreateDegree
func GetCmdCreateDegree(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createDegree [student] [subject] [batch]",
		Short: "Creates a new degree record",
		Args:  cobra.ExactArgs(3), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			batch, err := strconv.ParseUint(args[2], 16, 16)
			msg := types.NewMsgCreateDegree(cliCtx.GetFromAddress(), sdk.AccAddress(args[0]), args[1], uint16(batch))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
