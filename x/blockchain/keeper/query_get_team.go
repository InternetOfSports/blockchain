package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTeam(goCtx context.Context, req *types.QueryGetTeamRequest) (*types.QueryGetTeamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetTeamResponse{}, nil
}
