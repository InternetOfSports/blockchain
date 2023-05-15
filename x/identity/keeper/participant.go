package keeper

import (
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParticipant set a specific participant in the store from its index
func (k Keeper) SetParticipant(ctx sdk.Context, participant types.Participant) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKeyPrefix))
	b := k.cdc.MustMarshal(&participant)
	store.Set(types.ParticipantKey(
		participant.Index,
	), b)
}

// GetParticipant returns a participant from its index
func (k Keeper) GetParticipant(
	ctx sdk.Context,
	index string,

) (val types.Participant, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKeyPrefix))

	b := store.Get(types.ParticipantKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveParticipant removes a participant from the store
func (k Keeper) RemoveParticipant(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKeyPrefix))
	store.Delete(types.ParticipantKey(
		index,
	))
}

// GetAllParticipant returns all participant
func (k Keeper) GetAllParticipant(ctx sdk.Context) (list []types.Participant) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Participant
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
