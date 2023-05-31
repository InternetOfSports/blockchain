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

func CmdInviteParticipantToJoinTeam() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invite-participant-to-join-team [team] [participant]",
		Short: "Broadcast message invite-participant-to-join-team",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTeam := args[0]
			argParticipant := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInviteParticipantToJoinTeam(
				clientCtx.GetFromAddress().String(),
				argTeam,
				argParticipant,
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
