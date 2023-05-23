package keeper_test

import (
	"testing"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateTeam(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	msgServer.RegisterParticipant(ctx, types.NewMsgRegisterParticipant("cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28", "Messi"))
	tests := []struct {
		name string
		msg  types.MsgCreateTeam
		err  error
		resp types.MsgCreateTeamResponse
	}{
		{name: "Create FC Messi", msg: *types.NewMsgCreateTeam("cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28", "FC Messi"), resp: types.MsgCreateTeamResponse{}},
		{name: "Create FC Ronaldo", msg: *types.NewMsgCreateTeam("cosmos16fwqqp6v42qhyrk4fddqdhum5jmlc59ntd5tzm", "FC Ronaldo"), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No participant")},
		{name: "Create already existing FC Messi", msg: *types.NewMsgCreateTeam("cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28", "FC Messi"), err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team already exists")},
	}
	for _, tt := range tests {
		resp, respErr := msgServer.CreateTeam(ctx, &tt.msg)
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

func TestAssignedTeamOwner(t *testing.T) {
	owner := "cosmos1ph4efkd9vtt7wef2yhfvj430klup9jqu8fux28"
	ownerName := "Messi"
	teamName := "FC Messi"
	k, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)
	msgServer.RegisterParticipant(wctx, types.NewMsgRegisterParticipant(owner, ownerName))
	msgServer.CreateTeam(wctx, types.NewMsgCreateTeam(owner, teamName))
	teamResponse, err := k.Team(wctx, &types.QueryGetTeamRequest{Index: teamName})
	if err != nil {
		require.NoError(t, err)
	}
	if teamResponse.Team.Owner.Address != owner {
		require.ElementsMatch(t, teamResponse.Team.Owner.Address, owner)
	}

	if teamResponse.Team.Owner.Nickname != ownerName {
		require.ElementsMatch(t, teamResponse.Team.Owner.Nickname, ownerName)
	}
}
