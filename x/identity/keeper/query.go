package keeper

import (
	"github.com/InternetOfSports/blockchain/x/identity/types"
)

var _ types.QueryServer = Keeper{}
