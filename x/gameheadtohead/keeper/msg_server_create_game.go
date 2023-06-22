package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/gameheadtohead/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx

	return &types.MsgCreateGameResponse{}, nil
}
