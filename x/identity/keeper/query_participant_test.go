package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/testutil/nullify"
	"github.com/InternetOfSports/blockchain/x/identity/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestParticipantQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNParticipant(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetParticipantRequest
		response *types.QueryGetParticipantResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetParticipantRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetParticipantResponse{Participant: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetParticipantRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetParticipantResponse{Participant: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetParticipantRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Participant(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestParticipantQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNParticipant(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllParticipantRequest {
		return &types.QueryAllParticipantRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ParticipantAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Participant), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Participant),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ParticipantAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Participant), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Participant),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ParticipantAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Participant),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ParticipantAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
