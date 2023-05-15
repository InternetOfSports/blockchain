package keeper_test

import (
	"testing"

	testkeeper "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IdentityKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
