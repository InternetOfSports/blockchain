package keeper_test

import (
	"fmt"
	"testing"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestManageJoinTeamInvite(t *testing.T) {
	participant_messi := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	participant_ronaldo := "cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm"
	participant_ronaldinho := "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae"
	nonparticipant_davids := "cosmos1khr8vrjngr49rmdpzq0pg7rz037w9r3ew94r3p"
	team_messi := "FC Messi"

	tests := []struct {
		name string
		msg  types.MsgManageJoinTeamInvite
		err  error
		resp types.MsgManageJoinTeamInviteResponse
	}{
		{
			name: "Ronaldo accepts invite from Messi",
			msg:  *types.NewMsgManageJoinTeamInvite(participant_ronaldo, team_messi, true),
			resp: *&types.MsgManageJoinTeamInviteResponse{
				Team: &types.Team{
					Players: map[string]*types.Participant{
						participant_ronaldo: {
							Nickname: "Ronaldo",
							Address:  participant_ronaldo,
						},
					},
				},
			},
		},
		{
			name: "Ronaldo declines invite from Messi",
			msg:  *types.NewMsgManageJoinTeamInvite(participant_ronaldo, team_messi, false),
			resp: *&types.MsgManageJoinTeamInviteResponse{
				Team: &types.Team{},
			},
		},
		{
			name: "Ronaldo accepts invite from wrong team",
			msg:  *types.NewMsgManageJoinTeamInvite(participant_ronaldo, "FC Roni", true),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found"),
		},
		{
			name: "Ronaldinho accepts invite without invite",
			msg:  *types.NewMsgManageJoinTeamInvite(participant_ronaldinho, team_messi, true),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not invited"),
		},
		{
			name: "Non participants try to accept the invite",
			msg:  *types.NewMsgManageJoinTeamInvite(nonparticipant_davids, team_messi, true),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Participant not found"),
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		msgServer, ctx := setupMsgServer(t)
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_messi, "Messi"))
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldo, "Ronaldo"))
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldinho, "Ronaldinho"))
		msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_messi, team_messi))
		msgServer.SetTeamManager(ctx, types.NewMsgSetTeamManager(participant_messi, team_messi, participant_messi))
		msgServer.InviteParticipantToJoinTeam(ctx, types.NewMsgInviteParticipantToJoinTeam(participant_messi, team_messi, participant_ronaldo))
		resp, respErr := msgServer.ManageJoinTeamInvite(ctx, &tt.msg)

		err := tt.msg.ValidateBasic()
		require.NoError(t, err)

		if tt.err != nil || resp == nil {
			require.ErrorIs(t, respErr, tt.err)
			require.Equal(t, tt.err.Error(), respErr.Error())
		} else {
			require.NotNil(t, resp.Team)
			require.Equal(t, len(tt.resp.Team.Players), len(resp.Team.Players))
		}
	}
}
