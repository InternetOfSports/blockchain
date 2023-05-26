package keeper_test

import (
	"fmt"
	"testing"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestManageJoinTeamRequest(t *testing.T) {
	participant_messi := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	participant_ronaldo := "cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm"
	team_messi := "FC Messi"

	tests := []struct {
		name string
		msg  types.MsgManageJoinTeamRequest
		err  error
		resp types.MsgManageJoinTeamRequestResponse
	}{
		{
			name: "Allow ronaldo to join the team",
			msg:  *types.NewMsgManageJoinTeamRequest(participant_messi, team_messi, participant_ronaldo, true),
			resp: types.MsgManageJoinTeamRequestResponse{
				Team: &types.Team{
					Name: team_messi,
					Players: map[string]*types.Participant{
						participant_ronaldo: {
							Nickname: "ronaldo",
							Address:  participant_ronaldo,
						},
					},
					JoiningRequests: map[string]*types.Participant{},
				},
			},
		},
		{
			name: "Non existing Team",
			msg:  *types.NewMsgManageJoinTeamRequest(participant_messi, "Team X", participant_ronaldo, true),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found"),
		},
		{
			name: "No joining request",
			msg:  *types.NewMsgManageJoinTeamRequest(participant_messi, team_messi, participant_messi, true),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("No request from participant %s", participant_messi)),
		},
		{
			name: "Only team manager is allowed to accept",
			msg:  *types.NewMsgManageJoinTeamRequest(participant_ronaldo, team_messi, participant_ronaldo, true),
			err:  sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Only team manager can accept requests"),
		},
		{
			name: "Don't allow ronaldo to join the team",
			msg:  *types.NewMsgManageJoinTeamRequest(participant_messi, team_messi, participant_ronaldo, false),
			resp: types.MsgManageJoinTeamRequestResponse{
				Team: &types.Team{
					Name:            team_messi,
					Players:         map[string]*types.Participant{},
					JoiningRequests: map[string]*types.Participant{},
				},
			},
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.name)
		msgServer, ctx := setupMsgServer(t)
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_messi, "messi"))
		msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldo, "ronaldo"))
		msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_ronaldo, team_messi))
		msgServer.SetTeamManager(ctx, types.NewMsgSetTeamManager(participant_ronaldo, team_messi, participant_messi))
		msgServer.RequestJoinTeam(ctx, types.NewMsgRequestJoinTeam(participant_ronaldo, team_messi))
		resp, respErr := msgServer.ManageJoinTeamRequest(ctx, &tt.msg)

		err := tt.msg.ValidateBasic()
		require.NoError(t, err)

		if tt.err != nil || resp == nil {
			require.ErrorIs(t, respErr, tt.err)
			continue
		} else {
			require.Equal(t, len(tt.resp.Team.Players), len(resp.Team.Players))
		}

	}
}
