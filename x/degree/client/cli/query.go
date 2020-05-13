package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/akigugale/alt-verify/x/degree/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group degree queries under a subcommand
	degreeQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	degreeQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdListDegrees(queryRoute, cdc),
			GetCmdGetDegree(queryRoute, cdc),
			GetCmdGetDegreesOfUni(queryRoute, cdc),
		)...,
	)

	return degreeQueryCmd
}


func GetCmdListDegrees(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Short: "list",
		// Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListDegrees, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get degrees\n%s\n", err.Error())
				return nil
			}

			var out types.QueryResDegrees
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}


func GetCmdGetDegree(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "get [student id]",
		Short: "Query degree by student id (accAddress)",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			student := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetDegree, student), nil)
			if err != nil {
				fmt.Printf("could not resolve degree %s \n%s\n", student, err.Error())
				return nil
			}

			var out types.Degree
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}



func GetCmdGetDegreesOfUni(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "listofuni",
		Short: "List degrees of a particular uni.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			uni := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListDegrees, uni), nil)
			if err != nil {
				fmt.Printf("could not get uni %s degrees \n%s\n", uni, err.Error())
				return nil
			}

			var out types.QueryResDegrees
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}