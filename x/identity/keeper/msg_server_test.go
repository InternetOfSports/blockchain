package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/InternetOfSports/blockchain/testutil/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
