package gameheadtohead_test

import (
	"testing"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/testutil/nullify"
	"github.com/InternetOfSports/blockchain/x/gameheadtohead"
	"github.com/InternetOfSports/blockchain/x/gameheadtohead/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GameList: []types.Game{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GameheadtoheadKeeper(t)
	gameheadtohead.InitGenesis(ctx, *k, genesisState)
	got := gameheadtohead.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GameList, got.GameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
