package simulation

import (
	"math/rand"

	"github.com/InternetOfSports/blockchain/x/identity/keeper"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRegisterParticipant(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterParticipant{
			Creator:  simAccount.Address.String(),
			Nickname: "Messi",
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterParticipant simulation not implemented"), nil, nil
	}
}
