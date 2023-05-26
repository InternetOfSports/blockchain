package keeper

import (
	"context"
	"fmt"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ManageJoinTeamRequest(goCtx context.Context, msg *types.MsgManageJoinTeamRequest) (*types.MsgManageJoinTeamRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	team, teamFound := k.GetTeam(ctx, msg.Team)
	if !teamFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")
	}
	if team.Manager == nil || team.Manager.Address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Only team manager can accept/decline requests")
	}
	participant, participantJoiningRequestFound := team.JoiningRequests[msg.Participant]
	if !participantJoiningRequestFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("No request from participant %s", msg.Participant))
	}
	if msg.Accept {
		if team.Players == nil {
			team.Players = make(map[string]*types.Participant)
		}
		team.Players[participant.Address] = participant
	}
	delete(team.JoiningRequests, msg.Participant)
	k.SetTeam(ctx, team)
	return &types.MsgManageJoinTeamRequestResponse{Team: &team}, nil
}
