package gameheadtohead

import (
	"github.com/InternetOfSports/blockchain/x/gameheadtohead/keeper"
	"github.com/InternetOfSports/blockchain/x/gameheadtohead/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the game
	for _, elem := range genState.GameList {
		k.SetGame(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.GameList = k.GetAllGame(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
