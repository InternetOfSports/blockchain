package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RequestJoinTeam(goCtx context.Context, msg *types.MsgRequestJoinTeam) (*types.MsgRequestJoinTeamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	participant, participantFound := k.GetParticipant(ctx, msg.Creator)
	if !participantFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not a participant")
	}
	team, teamFound := k.GetTeam(ctx, msg.Team)
	if !teamFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")
	}
	if team.JoiningRequests == nil {
		team.JoiningRequests = make(map[string]*types.Participant)
	}
	team.JoiningRequests[participant.Address] = &participant
	k.SetTeam(ctx, team)
	return &types.MsgRequestJoinTeamResponse{}, nil
}
