package cli

import (
	"strconv"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdManageJoinTeamInvite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manage-join-team-invite [team] [accept]",
		Short: "Broadcast message manage-join-team-invite",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTeam := args[0]
			argAccept, err := cast.ToBoolE(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgManageJoinTeamInvite(
				clientCtx.GetFromAddress().String(),
				argTeam,
				argAccept,
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
