package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/faddat/clayfour/x/clayfour/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	clayfourTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	clayfourTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding # 1
		GetCmdCreatePost(cdc),
		GetCmdSetPost(cdc),
		GetCmdDeletePost(cdc),
	)...)

	return clayfourTxCmd
}
