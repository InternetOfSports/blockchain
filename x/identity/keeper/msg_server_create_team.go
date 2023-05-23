package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTeam(goCtx context.Context, msg *types.MsgCreateTeam) (*types.MsgCreateTeamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	participant, found := k.GetParticipant(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No participant")
	}
	_, foundTeam := k.GetTeam(ctx, msg.Name)
	if foundTeam {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team already exists")
	}
	k.SetTeam(ctx, types.Team{
		Index: msg.Name,
		Name:  msg.Name,
		Owner: &participant,
	})
	return &types.MsgCreateTeamResponse{}, nil
}
