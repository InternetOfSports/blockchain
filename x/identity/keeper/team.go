package keeper

import (
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTeam set a specific team in the store from its index
func (k Keeper) SetTeam(ctx sdk.Context, team types.Team) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamKeyPrefix))
	b := k.cdc.MustMarshal(&team)
	store.Set(types.TeamKey(
		team.Index,
	), b)
}

// GetTeam returns a team from its index
func (k Keeper) GetTeam(
	ctx sdk.Context,
	index string,

) (val types.Team, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamKeyPrefix))

	b := store.Get(types.TeamKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTeam removes a team from the store
func (k Keeper) RemoveTeam(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamKeyPrefix))
	store.Delete(types.TeamKey(
		index,
	))
}

// GetAllTeam returns all team
func (k Keeper) GetAllTeam(ctx sdk.Context) (list []types.Team) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Team
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
