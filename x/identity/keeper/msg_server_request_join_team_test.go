package keeper_test

import (
	"testing"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestRequestJoinTeam(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	team_fcMessi := "Fc Messi"
	participant_messi := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	participant_ronaldo := "cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm"
	no_participant_ronaldinho := "cosmos1hj78v8zjhmvp7jhlkvn4zjzxz5hujs75q4v3td"
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_messi, "Messi"))
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldo, "Ronaldo"))
	msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_messi, team_fcMessi))

	tests := []struct {
		name string
		msg  types.MsgRequestJoinTeam
		err  error
		resp types.MsgRequestJoinTeamResponse
	}{
		{name: "Ronaldo requests as player", msg: *types.NewMsgRequestJoinTeam(participant_ronaldo, team_fcMessi), resp: types.MsgRequestJoinTeamResponse{}},
		{name: "Ronaldinho requests as player", msg: *types.NewMsgRequestJoinTeam(no_participant_ronaldinho, team_fcMessi), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not a participant")},
		{name: "Messi requests as player to non existing Team", msg: *types.NewMsgRequestJoinTeam(participant_messi, "Fc Ronaldo"), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")},
	}
	for _, tt := range tests {
		resp, respErr := msgServer.RequestJoinTeam(ctx, &tt.msg)
		if resp != &tt.resp {
			require.ElementsMatch(t, resp, tt.resp)
		}
		if respErr != tt.err {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, respErr, tt.err)
				continue
			}
			require.NoError(t, err)
		}
		if resp != &tt.resp {
			require.ElementsMatch(t, resp, tt.resp)
		}
	}
}
