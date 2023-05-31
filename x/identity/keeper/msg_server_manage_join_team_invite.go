package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ManageJoinTeamInvite(goCtx context.Context, msg *types.MsgManageJoinTeamInvite) (*types.MsgManageJoinTeamInviteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	team, teamFound := k.GetTeam(ctx, msg.Team)
	if !teamFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")
	}
	participant, participantFound := k.GetParticipant(ctx, msg.Creator)
	if !participantFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Participant not found")
	}
	_, participantIsInvited := team.JoiningInvites[participant.Address]
	if !participantIsInvited {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not invited")
	}
	if msg.Accept {
		if team.Players == nil {
			team.Players = make(map[string]*types.Participant)
		}
		team.Players[participant.Address] = &participant
	}
	delete(team.JoiningInvites, participant.Address)
	delete(participant.TeamJoiningInvites, team.Name)
	k.SetTeam(ctx, team)
	k.SetParticipant(ctx, participant)
	return &types.MsgManageJoinTeamInviteResponse{Team: &team}, nil
}
