package cli

import (
	"strconv"

	"encoding/json"
	"github.com/bonedaddy/shitchain/x/shitstore/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBroadcastPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "broadcast-post [post]",
		Short: "Broadcast message broadcast-post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPost := new(types.Post)
			err = json.Unmarshal([]byte(args[0]), argPost)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBroadcastPost(
				clientCtx.GetFromAddress().String(),
				argPost,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
