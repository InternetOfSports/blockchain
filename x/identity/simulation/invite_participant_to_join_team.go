package simulation

import (
	"math/rand"

	"github.com/InternetOfSports/blockchain/x/identity/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgInviteParticipantToJoinTeam(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgInviteParticipantToJoinTeam{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the InviteParticipantToJoinTeam simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "InviteParticipantToJoinTeam simulation not implemented"), nil, nil
	}
}
