package keeper_test

import (
	"fmt"
	"testing"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestInviteParticipantToJoinTeam(t *testing.T) {
	participant_messi := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	participant_ronaldo := "cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm"
	team_messi := "FC Messi"

	tests := []struct {
		name string
		msg  types.MsgInviteParticipantToJoinTeam
		err  error
		resp types.MsgInviteParticipantToJoinTeamResponse
	}{
		{
			name: "Messi invites Ronaldo",
			msg:  *types.NewMsgInviteParticipantToJoinTeam(participant_messi, team_messi, participant_ronaldo),
			resp: *&types.MsgInviteParticipantToJoinTeamResponse{
				Participant: &types.Participant{
					Address:  participant_ronaldo,
					Nickname: "Ronaldo",
					TeamJoiningInvites: []string{
						team_messi,
					},
				},
			},
		},
		{
			name: "Messi invites Ronaldo to non existing team",
			msg:  *types.NewMsgInviteParticipantToJoinTeam(participant_messi, "Team X", participant_ronaldo),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found"),
		},
		{
			name: "Ronaldo tries to invite himself",
			msg:  *types.NewMsgInviteParticipantToJoinTeam(participant_ronaldo, team_messi, participant_ronaldo),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Only team manager can send invites"),
		},
		{
			name: "Messi invites non participant",
			msg:  *types.NewMsgInviteParticipantToJoinTeam(participant_messi, team_messi, "nobody"),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Participant not found"),
		},
	}

	for _, tt := range tests {
		fmt.Println(tt.name)
		msgServer, ctx := setupMsgServer(t)
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_messi, "Messi"))
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldo, "Ronaldo"))
		msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_messi, team_messi))
		msgServer.SetTeamManager(ctx, types.NewMsgSetTeamManager(participant_messi, team_messi, participant_messi))
		resp, respErr := msgServer.InviteParticipantToJoinTeam(ctx, &tt.msg)

		err := tt.msg.ValidateBasic()
		require.NoError(t, err)

		if tt.err != nil || resp == nil {
			require.ErrorIs(t, respErr, tt.err)
			require.Equal(t, respErr.Error(), tt.err.Error())
		} else {
			require.NotNil(t, resp.Participant)
			require.Equal(t, len(tt.resp.Participant.TeamJoiningInvites), len(resp.Participant.TeamJoiningInvites))
		}
	}

}

func TestMultipleInviteParticipantToJoinTeam(t *testing.T) {
	participant_messi := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	participant_ronaldo := "cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm"
	team_messi := "FC Messi"
	team_messi02 := "FC Messi 02"

	msgServer, ctx := setupMsgServer(t)
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_messi, "Messi"))
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldo, "Ronaldo"))
	msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_messi, team_messi))
	msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_messi, team_messi02))
	msgServer.SetTeamManager(ctx, types.NewMsgSetTeamManager(participant_messi, team_messi, participant_messi))
	msgServer.SetTeamManager(ctx, types.NewMsgSetTeamManager(participant_messi, team_messi02, participant_messi))
	_, respErr := msgServer.InviteParticipantToJoinTeam(ctx, types.NewMsgInviteParticipantToJoinTeam(participant_messi, team_messi, participant_messi))
	_, respErr = msgServer.InviteParticipantToJoinTeam(ctx, types.NewMsgInviteParticipantToJoinTeam(participant_messi, team_messi02, participant_messi))
	require.NoError(t, respErr)
}
