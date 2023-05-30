package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) InviteParticipantToJoinTeam(goCtx context.Context, msg *types.MsgInviteParticipantToJoinTeam) (*types.MsgInviteParticipantToJoinTeamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	team, teamFound := k.GetTeam(ctx, msg.Team)
	if !teamFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")
	}
	if team.Manager == nil || team.Manager.Address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Only team manager can send invites")
	}
	participant, participantFound := k.GetParticipant(ctx, msg.Participant)
	if !participantFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Participant not found")
	}
	if team.JoiningInvites == nil {
		team.JoiningInvites = make(map[string]*types.Participant)
	}
	team.JoiningInvites[participant.Address] = &participant
	if participant.TeamJoiningInvites == nil {
		participant.TeamJoiningInvites = []string{
			team.Name,
		}
	} else {
		participant.TeamJoiningInvites = append(participant.TeamJoiningInvites, team.Name)
	}
	k.SetTeam(ctx, team)
	k.SetParticipant(ctx, participant)
	return &types.MsgInviteParticipantToJoinTeamResponse{Participant: &participant}, nil
}
