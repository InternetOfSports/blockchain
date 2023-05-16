package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterParticipant(goCtx context.Context, msg *types.MsgRegisterParticipant) (*types.MsgRegisterParticipantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetParticipant(ctx, types.Participant{
		Nickname: msg.Nickname,
		Address:  msg.Creator,
		Index:    msg.Creator,
	})

	return &types.MsgRegisterParticipantResponse{}, nil
}
