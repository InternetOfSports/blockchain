package keeper

import (
	"context"

	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TeamAll(goCtx context.Context, req *types.QueryAllTeamRequest) (*types.QueryAllTeamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var teams []types.Team
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	teamStore := prefix.NewStore(store, types.KeyPrefix(types.TeamKeyPrefix))

	pageRes, err := query.Paginate(teamStore, req.Pagination, func(key []byte, value []byte) error {
		var team types.Team
		if err := k.cdc.Unmarshal(value, &team); err != nil {
			return err
		}

		teams = append(teams, team)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTeamResponse{Team: teams, Pagination: pageRes}, nil
}

func (k Keeper) Team(goCtx context.Context, req *types.QueryGetTeamRequest) (*types.QueryGetTeamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTeam(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTeamResponse{Team: val}, nil
}
