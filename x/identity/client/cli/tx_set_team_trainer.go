package cli

import (
	"strconv"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetTeamTrainer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-team-trainer [team] [trainer]",
		Short: "Broadcast message set-team-trainer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTeam := args[0]
			argTrainer := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetTeamTrainer(
				clientCtx.GetFromAddress().String(),
				argTeam,
				argTrainer,
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
