package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetTeamTrainer(goCtx context.Context, msg *types.MsgSetTeamTrainer) (*types.MsgSetTeamTrainerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	participant, participantFound := k.GetParticipant(ctx, msg.Trainer)
	if !participantFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Trainer not found (not a participant)")
	}
	team, found := k.GetTeam(ctx, msg.Team)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Team not found")
	}
	if team.Owner.Address != msg.Creator || (team.Manager != nil && team.Manager.Address != msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Not allowed to change the trainer")
	}
	team.Trainer = &participant
	if participant.TeamsTraining == nil {
		participant.TeamsTraining = make(map[string]string)
	}
	participant.TeamsTraining[team.Name] = team.Name
	k.SetTeam(ctx, team)
	return &types.MsgSetTeamTrainerResponse{}, nil
}
