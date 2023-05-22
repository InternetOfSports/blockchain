package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RegisterParticipant(goCtx context.Context, msg *types.MsgRegisterParticipant) (*types.MsgRegisterParticipantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, foundAddress := k.GetParticipant(ctx, msg.Creator)
	_, foundNickname := k.GetParticipant(ctx, msg.Nickname)

	if foundAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Address already participated")
	}
	if foundNickname {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Nickname already exists")
	}

	k.SetParticipant(ctx, types.Participant{
		Nickname: msg.Nickname,
		Address:  msg.Creator,
		Index:    msg.Creator,
	})
	k.SetParticipant(ctx, types.Participant{
		Nickname: msg.Nickname,
		Address:  msg.Creator,
		Index:    msg.Nickname,
	})

	return &types.MsgRegisterParticipantResponse{}, nil
}
