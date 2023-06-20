package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetTeamManager(goCtx context.Context, msg *types.MsgSetTeamManager) (*types.MsgSetTeamManagerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	participant, participantFound := k.GetParticipant(ctx, msg.Manager)
	if !participantFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Manager not found (not a participant)")
	}
	team, found := k.GetTeam(ctx, msg.Team)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")
	}
	if team.Owner.Address != msg.Creator || (team.Manager != nil && team.Manager.Address != msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not allowed to change the manager")
	}
	if participant.TeamsManaging == nil {
		participant.TeamsManaging = make(map[string]string)
	}
	participant.TeamsManaging[team.Name] = team.Name
	team.Manager = &participant
	k.SetTeam(ctx, team)

	return &types.MsgSetTeamManagerResponse{}, nil
}
