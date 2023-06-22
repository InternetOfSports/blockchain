package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/testutil/nullify"
	"github.com/InternetOfSports/blockchain/x/gameheadtohead/keeper"
	"github.com/InternetOfSports/blockchain/x/gameheadtohead/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGame(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Game {
	items := make([]types.Game, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetGame(ctx, items[i])
	}
	return items
}

func TestGameGet(t *testing.T) {
	keeper, ctx := keepertest.GameheadtoheadKeeper(t)
	items := createNGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGameRemove(t *testing.T) {
	keeper, ctx := keepertest.GameheadtoheadKeeper(t)
	items := createNGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGame(ctx,
			item.Index,
		)
		_, found := keeper.GetGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.GameheadtoheadKeeper(t)
	items := createNGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGame(ctx)),
	)
}
