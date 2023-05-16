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

func createNTeam(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Team {
	items := make([]types.Team, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTeam(ctx, items[i])
	}
	return items
}

func TestTeamGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNTeam(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTeam(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTeamRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNTeam(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTeam(ctx,
			item.Index,
		)
		_, found := keeper.GetTeam(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTeamGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNTeam(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTeam(ctx)),
	)
}
