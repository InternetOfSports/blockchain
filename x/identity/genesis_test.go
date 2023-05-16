package identity_test

import (
	"testing"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/testutil/nullify"
	"github.com/InternetOfSports/blockchain/x/identity"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ParticipantList: []types.Participant{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		TeamList: []types.Team{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ParticipantList, got.ParticipantList)
	require.ElementsMatch(t, genesisState.TeamList, got.TeamList)
	// this line is used by starport scaffolding # genesis/test/assert
}
