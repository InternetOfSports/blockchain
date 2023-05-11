package keeper

import (
	"github.com/InternetOfSports/blockchain/x/blockchain/types"
)

var _ types.QueryServer = Keeper{}
