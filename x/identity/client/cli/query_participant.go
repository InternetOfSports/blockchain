package cli

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListParticipant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-participant",
		Short: "list all Participant",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllParticipantRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ParticipantAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowParticipant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-participant [index]",
		Short: "shows a Participant",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetParticipantRequest{
				Index: argIndex,
			}

			res, err := queryClient.Participant(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
