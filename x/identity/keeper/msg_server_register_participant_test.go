package keeper_test

import (
	"testing"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestRegisterParticipant(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	tests := []struct {
		name string
		msg  types.MsgRegisterParticipant
		err  error
		resp types.MsgRegisterParticipantResponse
	}{
		{name: "Add Alice", msg: *types.NewMsgRegisterParticipant("cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28", "Alice"), resp: types.MsgRegisterParticipantResponse{}},
		{name: "Add Bob", msg: *types.NewMsgRegisterParticipant("cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28", "Bob"), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Address already participated")},
		{name: "Add Alice", msg: *types.NewMsgRegisterParticipant("cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm", "Alice"), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Nickname already exists")},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		resp, respErr := msgServer.RegisterParticipant(ctx, &tt.msg)
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
