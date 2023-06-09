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

func (k Keeper) ParticipantAll(goCtx context.Context, req *types.QueryAllParticipantRequest) (*types.QueryAllParticipantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var participants []types.Participant
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	participantStore := prefix.NewStore(store, types.KeyPrefix(types.ParticipantKeyPrefix))

	pageRes, err := query.Paginate(participantStore, req.Pagination, func(key []byte, value []byte) error {
		var participant types.Participant
		if err := k.cdc.Unmarshal(value, &participant); err != nil {
			return err
		}

		participants = append(participants, participant)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllParticipantResponse{Participant: participants, Pagination: pageRes}, nil
}

func (k Keeper) Participant(goCtx context.Context, req *types.QueryGetParticipantRequest) (*types.QueryGetParticipantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetParticipant(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetParticipantResponse{Participant: val}, nil
}
