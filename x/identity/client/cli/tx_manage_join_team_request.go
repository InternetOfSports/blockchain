package cli

import (
    "strconv"
	
	 "github.com/spf13/cast"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/InternetOfSports/blockchain/x/identity/types"
)

var _ = strconv.Itoa(0)

func CmdManageJoinTeamRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manage-join-team-request [team] [participant] [accept]",
		Short: "Broadcast message manage-join-team-request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argTeam := args[0]
             argParticipant := args[1]
             argAccept, err := cast.ToBoolE(args[2])
            		if err != nil {
                		return err
            		}
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgManageJoinTeamRequest(
				clientCtx.GetFromAddress().String(),
				argTeam,
				argParticipant,
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