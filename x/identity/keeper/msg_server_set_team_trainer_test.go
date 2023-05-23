package keeper_test

import (
	"testing"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestSetTeamTrainer(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	participant_messi := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	participant_ronaldo := "cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm"
	no_participant_ronaldinho := "cosmos1hj78v8zjhmvp7jhlkvn4zjzxz5hujs75q4v3td"
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_messi, "Messi"))
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant(participant_ronaldo, "Ronaldo"))
	msgServer.CreateTeam(ctx, types.NewMsgCreateTeam(participant_messi, "Fc Messi"))
	tests := []struct {
		name string
		msg  types.MsgSetTeamTrainer
		err  error
		resp types.MsgSetTeamTrainerResponse
	}{
		{name: "Set Messi as Trainer", msg: *types.NewMsgSetTeamTrainer(participant_messi, "Fc Messi", participant_messi), resp: types.MsgSetTeamTrainerResponse{}},
		{name: "Set Ronaldo as Trainer", msg: *types.NewMsgSetTeamTrainer(participant_messi, "Fc Messi", participant_ronaldo), resp: types.MsgSetTeamTrainerResponse{}},
		{name: "Set Ronaldinho as Trainer", msg: *types.NewMsgSetTeamTrainer(participant_messi, "Fc Messi", no_participant_ronaldinho), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Trainer not found (not a participant)")},
		{name: "Set Ronaldo as Trainer", msg: *types.NewMsgSetTeamTrainer(participant_ronaldo, "Fc Messi", participant_ronaldo), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not allowed to change the trainer")},
		{name: "Set Ronaldo as Trainer", msg: *types.NewMsgSetTeamTrainer(participant_ronaldo, "Fc Ronaldo", participant_ronaldo), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")},
	}
	for _, tt := range tests {
		resp, respErr := msgServer.SetTeamTrainer(ctx, &tt.msg)
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
