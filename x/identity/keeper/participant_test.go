package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/testutil/nullify"
	"github.com/InternetOfSports/blockchain/x/identity/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNParticipant(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Participant {
	items := make([]types.Participant, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetParticipant(ctx, items[i])
	}
	return items
}

func TestParticipantGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetParticipant(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestParticipantRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveParticipant(ctx,
			item.Index,
		)
		_, found := keeper.GetParticipant(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestParticipantGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllParticipant(ctx)),
	)
}
